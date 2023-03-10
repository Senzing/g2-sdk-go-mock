/*
 *
 */

// Package g2configmgrclient implements a client for the service.
package g2configmgr

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	g2configmgrapi "github.com/senzing/g2-sdk-go/g2configmgr"
	"github.com/senzing/go-logging/logger"
	"github.com/senzing/go-logging/messagelogger"
	"github.com/senzing/go-observing/observer"
	"github.com/senzing/go-observing/subject"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type G2configmgr struct {
	isTrace                  bool
	logger                   messagelogger.MessageLoggerInterface
	observers                subject.Subject
	AddConfigResult          int64
	GetConfigResult          string
	GetConfigListResult      string
	GetDefaultConfigIDResult int64
}

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

// Get the Logger singleton.
func (client *G2configmgr) getLogger() messagelogger.MessageLoggerInterface {
	if client.logger == nil {
		client.logger, _ = messagelogger.NewSenzingApiLogger(ProductId, g2configmgrapi.IdMessages, g2configmgrapi.IdStatuses, messagelogger.LevelInfo)
	}
	return client.logger
}

// Notify registered observers.
func (client *G2configmgr) notify(ctx context.Context, messageId int, err error, details map[string]string) {
	now := time.Now()
	details["subjectId"] = strconv.Itoa(ProductId)
	details["messageId"] = strconv.Itoa(messageId)
	details["messageTime"] = strconv.FormatInt(now.UnixNano(), 10)
	if err != nil {
		details["error"] = err.Error()
	}
	message, err := json.Marshal(details)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		client.observers.NotifyObservers(ctx, string(message))
	}
}

// Trace method entry.
func (client *G2configmgr) traceEntry(errorNumber int, details ...interface{}) {
	client.getLogger().Log(errorNumber, details...)
}

// Trace method exit.
func (client *G2configmgr) traceExit(errorNumber int, details ...interface{}) {
	client.getLogger().Log(errorNumber, details...)
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

/*
The AddConfig method adds a Senzing configuration JSON document to the Senzing database.

Input
  - ctx: A context to control lifecycle.
  - configStr: The Senzing configuration JSON document.
  - configComments: A free-form string of comments describing the configuration document.

Output
  - A configuration identifier.
*/
func (client *G2configmgr) AddConfig(ctx context.Context, configStr string, configComments string) (int64, error) {
	if client.isTrace {
		client.traceEntry(1, configStr, configComments)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"configComments": configComments,
			}
			client.notify(ctx, 8001, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(2, configStr, configComments, client.AddConfigResult, err, time.Since(entryTime))
	}
	return client.AddConfigResult, err
}

/*
The Destroy method will destroy and perform cleanup for the Senzing G2ConfigMgr object.
It should be called after all other calls are complete.

Input
  - ctx: A context to control lifecycle.
*/
func (client *G2configmgr) Destroy(ctx context.Context) error {
	if client.isTrace {
		client.traceEntry(5)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{}
			client.notify(ctx, 8002, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(6, err, time.Since(entryTime))
	}
	return err
}

/*
The GetConfig method retrieves a specific Senzing configuration JSON document from the Senzing database.

Input
  - ctx: A context to control lifecycle.
  - configID: The configuration identifier of the desired Senzing Engine configuration JSON document to retrieve.

Output
  - A JSON document containing the Senzing configuration.
    See the example output.
*/
func (client *G2configmgr) GetConfig(ctx context.Context, configID int64) (string, error) {
	if client.isTrace {
		client.traceEntry(7, configID)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{}
			client.notify(ctx, 8003, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(8, configID, client.GetConfigResult, err, time.Since(entryTime))
	}
	return client.GetConfigResult, err
}

/*
The GetConfigList method retrieves a list of Senzing configurations from the Senzing database.

Input
  - ctx: A context to control lifecycle.

Output
  - A JSON document containing Senzing configurations.
    See the example output.
*/
func (client *G2configmgr) GetConfigList(ctx context.Context) (string, error) {
	if client.isTrace {
		client.traceEntry(9)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{}
			client.notify(ctx, 8004, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(10, client.GetConfigListResult, err, time.Since(entryTime))
	}
	return client.GetConfigListResult, err
}

/*
The GetDefaultConfigID method retrieves from the Senzing database the configuration identifier of the default Senzing configuration.

Input
  - ctx: A context to control lifecycle.

Output
  - A configuration identifier which identifies the current configuration in use.
*/
func (client *G2configmgr) GetDefaultConfigID(ctx context.Context) (int64, error) {
	if client.isTrace {
		client.traceEntry(11)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{}
			client.notify(ctx, 8005, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(12, client.GetDefaultConfigIDResult, err, time.Since(entryTime))
	}
	return client.GetDefaultConfigIDResult, err
}

/*
The GetSdkId method returns the identifier of this particular Software Development Kit (SDK).
It is handy when working with multiple implementations of the same G2configmgrInterface.
For this implementation, "mock" is returned.

Input
  - ctx: A context to control lifecycle.
*/
func (client *G2configmgr) GetSdkId(ctx context.Context) string {
	if client.isTrace {
		client.traceEntry(29)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{}
			client.notify(ctx, 8010, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(30, err, time.Since(entryTime))
	}
	return "mock"
}

/*
The Init method initializes the Senzing G2ConfigMgr object.
It must be called prior to any other calls.

Input
  - ctx: A context to control lifecycle.
  - moduleName: A name for the auditing node, to help identify it within system logs.
  - iniParams: A JSON string containing configuration parameters.
  - verboseLogging: A flag to enable deeper logging of the G2 processing. 0 for no Senzing logging; 1 for logging.
*/
func (client *G2configmgr) Init(ctx context.Context, moduleName string, iniParams string, verboseLogging int) error {
	if client.isTrace {
		client.traceEntry(17, moduleName, iniParams, verboseLogging)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"iniParams":      iniParams,
				"moduleName":     moduleName,
				"verboseLogging": strconv.Itoa(verboseLogging),
			}
			client.notify(ctx, 8006, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(18, moduleName, iniParams, verboseLogging, err, time.Since(entryTime))
	}
	return err
}

/*
The RegisterObserver method adds the observer to the list of observers notified.

Input
  - ctx: A context to control lifecycle.
  - observer: The observer to be added.
*/
func (client *G2configmgr) RegisterObserver(ctx context.Context, observer observer.Observer) error {
	if client.isTrace {
		client.traceEntry(25, observer.GetObserverId(ctx))
	}
	entryTime := time.Now()
	if client.observers == nil {
		client.observers = &subject.SubjectImpl{}
	}
	err := client.observers.RegisterObserver(ctx, observer)
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"observerID": observer.GetObserverId(ctx),
			}
			client.notify(ctx, 8010, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(26, observer.GetObserverId(ctx), err, time.Since(entryTime))
	}
	return err
}

/*
The ReplaceDefaultConfigID method replaces the old configuration identifier with a new configuration identifier in the Senzing database.
It is like a "compare-and-swap" instruction to serialize concurrent editing of configuration.
If oldConfigID is no longer the "old configuration identifier", the operation will fail.
To simply set the default configuration ID, use SetDefaultConfigID().

Input
  - ctx: A context to control lifecycle.
  - oldConfigID: The configuration identifier to replace.
  - newConfigID: The configuration identifier to use as the default.
*/
func (client *G2configmgr) ReplaceDefaultConfigID(ctx context.Context, oldConfigID int64, newConfigID int64) error {
	if client.isTrace {
		client.traceEntry(19, oldConfigID, newConfigID)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"newConfigID": strconv.FormatInt(newConfigID, 10),
			}
			client.notify(ctx, 8007, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(20, oldConfigID, newConfigID, err, time.Since(entryTime))
	}
	return err
}

/*
The SetDefaultConfigID method replaces the sets a new configuration identifier in the Senzing database.
To serialize modifying of the configuration identifier, see ReplaceDefaultConfigID().

Input
  - ctx: A context to control lifecycle.
  - configID: The configuration identifier of the Senzing Engine configuration to use as the default.
*/
func (client *G2configmgr) SetDefaultConfigID(ctx context.Context, configID int64) error {
	if client.isTrace {
		client.traceEntry(21, configID)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"configID": strconv.FormatInt(configID, 10),
			}
			client.notify(ctx, 8008, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(22, configID, err, time.Since(entryTime))
	}
	return err
}

/*
The SetLogLevel method sets the level of logging.

Input
  - ctx: A context to control lifecycle.
  - logLevel: The desired log level. TRACE, DEBUG, INFO, WARN, ERROR, FATAL or PANIC.
*/
func (client *G2configmgr) SetLogLevel(ctx context.Context, logLevel logger.Level) error {
	if client.isTrace {
		client.traceEntry(23, logLevel)
	}
	entryTime := time.Now()
	var err error = nil
	client.getLogger().SetLogLevel(messagelogger.Level(logLevel))
	client.isTrace = (client.getLogger().GetLogLevel() == messagelogger.LevelTrace)
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"logLevel": logger.LevelToTextMap[logLevel],
			}
			client.notify(ctx, 8011, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(24, logLevel, err, time.Since(entryTime))
	}
	return err
}

/*
The UnregisterObserver method removes the observer to the list of observers notified.

Input
  - ctx: A context to control lifecycle.
  - observer: The observer to be added.
*/
func (client *G2configmgr) UnregisterObserver(ctx context.Context, observer observer.Observer) error {
	if client.isTrace {
		client.traceEntry(27, observer.GetObserverId(ctx))
	}
	entryTime := time.Now()
	var err error = nil
	if client.observers != nil {
		// Tricky code:
		// client.notify is called synchronously before client.observers is set to nil.
		// In client.notify, each observer will get notified in a goroutine.
		// Then client.observers may be set to nil, but observer goroutines will be OK.
		details := map[string]string{
			"observerID": observer.GetObserverId(ctx),
		}
		client.notify(ctx, 8012, err, details)
	}
	err = client.observers.UnregisterObserver(ctx, observer)
	if !client.observers.HasObservers(ctx) {
		client.observers = nil
	}
	if client.isTrace {
		defer client.traceExit(28, observer.GetObserverId(ctx), err, time.Since(entryTime))
	}
	return err
}
