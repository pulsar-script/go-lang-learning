package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// LogEntry represents a single log message.
type LogEntry struct {
	Timestamp time.Time
	Level     string
	Message   string
}

// Configuration for the application.
var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "log-monitor",
	Short: "A CLI tool to monitor logs in real-time.",
	Long: `log-monitor is a terminal-based application that provides a real-time
view of logs, including a summary and distribution of log levels.`,
	Run: func(cmd *cobra.Command, args []string) {
		startUI()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.log-monitor.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".log-monitor" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".log-monitor")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

// startUI initializes and runs the terminal user interface.
func startUI() {
	app := tview.NewApplication()

	// Create the main layout components.
	logTable := tview.NewTable().
		SetBorders(true).
		SetSelectable(true, false)
	logTable.SetTitle("Live Logs").SetBorder(true)

	summaryView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true)
	summaryView.SetTitle("Summary").SetBorder(true)

	chartView := tview.NewTextView().
		SetDynamicColors(true)
	chartView.SetTitle("Log Level Distribution").SetBorder(true)

	// Create a flex layout to arrange the components.
	flex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(summaryView, 0, 1, false).
			AddItem(chartView, 0, 1, false), 0, 1, false).
		AddItem(logTable, 0, 3, true)

	// Set up the log generation and UI updates.
	logChan := make(chan LogEntry)
	go generateLogs(logChan)
	go updateUI(app, logTable, summaryView, chartView, logChan)

	// Set the root of the application and run it.
	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		log.Fatalf("Error running application: %v", err)
	}
}

// generateLogs simulates log generation and sends them to a channel.
func generateLogs(logChan chan<- LogEntry) {
	logLevels := []string{"INFO", "WARN", "ERROR", "DEBUG"}
	messages := []string{
		"User logged in successfully",
		"Failed to connect to database",
		"Payment processed",
		"Request timed out",
		"New user registered",
		"Invalid input received",
		"Cache cleared",
		"API rate limit exceeded",
	}

	for {
		entry := LogEntry{
			Timestamp: time.Now(),
			Level:     logLevels[rand.Intn(len(logLevels))],
			Message:   messages[rand.Intn(len(messages))],
		}
		logChan <- entry
		time.Sleep(time.Duration(rand.Intn(1000)+200) * time.Millisecond)
	}
}

// updateUI receives logs and updates the TUI components.
func updateUI(app *tview.Application, table *tview.Table, summary *tview.TextView, chart *tview.TextView, logChan <-chan LogEntry) {
	var logs []LogEntry
	logCounts := make(map[string]int)

	// Initial table headers.
	table.SetCell(0, 0, tview.NewTableCell("Timestamp").SetSelectable(false).SetTextColor(tcell.ColorYellow))
	table.SetCell(0, 1, tview.NewTableCell("Level").SetSelectable(false).SetTextColor(tcell.ColorYellow))
	table.SetCell(0, 2, tview.NewTableCell("Message").SetSelectable(false).SetTextColor(tcell.ColorYellow))

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case entry := <-logChan:
			logs = append(logs, entry)
			logCounts[entry.Level]++

			// Add new row to the table.
			rowCount := table.GetRowCount()
			table.SetCell(rowCount, 0, tview.NewTableCell(entry.Timestamp.Format("15:04:05")).SetTextColor(tcell.ColorDarkGray))
			table.SetCell(rowCount, 1, tview.NewTableCell(entry.Level).SetTextColor(getLevelColor(entry.Level)))
			table.SetCell(rowCount, 2, tview.NewTableCell(entry.Message))
			table.ScrollToEnd()

		case <-ticker.C:
			// Update summary view.
			summary.Clear()
			fmt.Fprintf(summary, "Total Logs: [yellow]%d\n", len(logs))
			fmt.Fprintf(summary, "Errors: [red]%d\n", logCounts["ERROR"])
			fmt.Fprintf(summary, "Warnings: [orange]%d", logCounts["WARN"])

			// Update chart view.
			chart.Clear()
			total := len(logs)
			if total > 0 {
				for _, level := range []string{"INFO", "DEBUG", "WARN", "ERROR"} {
					count := logCounts[level]
					percentage := float64(count) / float64(total) * 100
					bar := createBar(int(percentage/2), getLevelColor(level))
					fmt.Fprintf(chart, "%-5s: [white]%s %.2f%%\n", level, bar, percentage)
				}
			}

			// Redraw the application.
			app.Draw()
		}
	}
}

// getLevelColor returns a tcell.Color based on the log level.
func getLevelColor(level string) tcell.Color {
	switch level {
	case "INFO":
		return tcell.ColorGreen
	case "WARN":
		return tcell.ColorOrange
	case "ERROR":
		return tcell.ColorRed
	case "DEBUG":
		return tcell.ColorBlue
	default:
		return tcell.ColorWhite
	}
}

// createBar generates a simple text-based bar for the chart.
func createBar(length int, color tcell.Color) string {
	bar := ""
	for i := 0; i < length; i++ {
		bar += "█"
	}
	return fmt.Sprintf("[%s]%s", color.String(), bar)
}

func main() {
	Execute()
}
