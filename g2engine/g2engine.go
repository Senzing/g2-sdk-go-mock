/*
 *
 */

// Package g2engine implements a client for the service.
package g2engine

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	g2engineapi "github.com/senzing/g2-sdk-go/g2engine"
	"github.com/senzing/go-logging/logger"
	"github.com/senzing/go-logging/messagelogger"
	"github.com/senzing/go-observing/observer"
	"github.com/senzing/go-observing/subject"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type G2engine struct {
	isTrace                                                bool
	logger                                                 messagelogger.MessageLoggerInterface
	observers                                              subject.Subject
	AddRecordWithInfoResult                                string
	AddRecordWithInfoWithReturnedRecordIDResultGetWithInfo string
	AddRecordWithInfoWithReturnedRecordIDResultRecordID    string
	AddRecordWithReturnedRecordIDResult                    string
	CheckRecordResult                                      string
	CountRedoRecordsResult                                 int64
	DeleteRecordWithInfoResult                             string
	ExportConfigResult                                     string
	ExportConfigAndConfigIDResultConfig                    string
	ExportConfigAndConfigIDResultConfigID                  int64
	ExportCSVEntityReportResult                            uintptr
	ExportJSONEntityReportResult                           uintptr
	FetchNextResult                                        string
	FindInterestingEntitiesByEntityIDResult                string
	FindInterestingEntitiesByRecordIDResult                string
	FindNetworkByEntityID_V2Result                         string
	FindNetworkByEntityIDResult                            string
	FindNetworkByRecordID_V2Result                         string
	FindNetworkByRecordIDResult                            string
	FindPathByEntityID_V2Result                            string
	FindPathByEntityIDResult                               string
	FindPathByRecordID_V2Result                            string
	FindPathByRecordIDResult                               string
	FindPathExcludingByEntityID_V2Result                   string
	FindPathExcludingByEntityIDResult                      string
	FindPathExcludingByRecordID_V2Result                   string
	FindPathExcludingByRecordIDResult                      string
	FindPathIncludingSourceByEntityID_V2Result             string
	FindPathIncludingSourceByEntityIDResult                string
	FindPathIncludingSourceByRecordID_V2Result             string
	FindPathIncludingSourceByRecordIDResult                string
	GetActiveConfigIDResult                                int64
	GetEntityByEntityID_V2Result                           string
	GetEntityByEntityIDResult                              string
	GetEntityByRecordID_V2Result                           string
	GetEntityByRecordIDResult                              string
	GetRecord_V2Result                                     string
	GetRecordResult                                        string
	GetRedoRecordResult                                    string
	GetRepositoryLastModifiedTimeResult                    int64
	GetVirtualEntityByRecordID_V2Result                    string
	GetVirtualEntityByRecordIDResult                       string
	HowEntityByEntityID_V2Result                           string
	HowEntityByEntityIDResult                              string
	ProcessRedoRecordResult                                string
	ProcessRedoRecordWithInfoResult                        string
	ProcessRedoRecordWithInfoResultWithInfo                string
	ProcessWithInfoResult                                  string
	ProcessWithResponseResult                              string
	ProcessWithResponseResizeResult                        string
	ReevaluateEntityWithInfoResult                         string
	ReevaluateRecordWithInfoResult                         string
	ReplaceRecordWithInfoResult                            string
	SearchByAttributes_V2Result                            string
	SearchByAttributesResult                               string
	StatsResult                                            string
	WhyEntities_V2Result                                   string
	WhyEntitiesResult                                      string
	WhyEntityByEntityID_V2Result                           string
	WhyEntityByEntityIDResult                              string
	WhyEntityByRecordID_V2Result                           string
	WhyEntityByRecordIDResult                              string
	WhyRecords_V2Result                                    string
	WhyRecordsResult                                       string
}

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

// Get the Logger singleton.
func (client *G2engine) getLogger() messagelogger.MessageLoggerInterface {
	if client.logger == nil {
		client.logger, _ = messagelogger.NewSenzingApiLogger(ProductId, g2engineapi.IdMessages, g2engineapi.IdStatuses, messagelogger.LevelInfo)
	}
	return client.logger
}

// Notify registered observers.
func (client *G2engine) notify(ctx context.Context, messageId int, err error, details map[string]string) {
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
func (client *G2engine) traceEntry(errorNumber int, details ...interface{}) {
	client.getLogger().Log(errorNumber, details...)
}

// Trace method exit.
func (client *G2engine) traceExit(errorNumber int, details ...interface{}) {
	client.getLogger().Log(errorNumber, details...)
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

/*
The AddRecord method adds a record into the Senzing repository.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - recordID: The unique identifier within the records of the same data source.
  - jsonData: A JSON document containing the record to be added to the Senzing repository.
  - loadID: An identifier used to distinguish different load batches/sessions. An empty string is acceptable.
*/
func (client *G2engine) AddRecord(ctx context.Context, dataSourceCode string, recordID string, jsonData string, loadID string) error {
	if client.isTrace {
		client.traceEntry(1, dataSourceCode, recordID, jsonData, loadID)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"dataSourceCode": dataSourceCode,
				"recordID":       recordID,
				"loadID":         loadID,
			}
			client.notify(ctx, 8001, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(2, dataSourceCode, recordID, jsonData, loadID, err, time.Since(entryTime))
	}
	return err
}

/*
The AddRecordWithInfo method adds a record into the Senzing repository and returns information on the affected entities.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - recordID: The unique identifier within the records of the same data source.
  - jsonData: A JSON document containing the record to be added to the Senzing repository.
  - loadID: An identifier used to distinguish different load batches/sessions. An empty string is acceptable.
  - flags: Flags used to control information returned.

Output
  - A JSON document.
    See the example output.
*/
func (client *G2engine) AddRecordWithInfo(ctx context.Context, dataSourceCode string, recordID string, jsonData string, loadID string, flags int64) (string, error) {
	if client.isTrace {
		client.traceEntry(3, dataSourceCode, recordID, jsonData, loadID, flags)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"dataSourceCode": dataSourceCode,
				"recordID":       recordID,
				"loadID":         loadID,
			}
			client.notify(ctx, 8002, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(4, dataSourceCode, recordID, jsonData, loadID, flags, client.AddRecordWithInfoResult, err, time.Since(entryTime))
	}
	return client.AddRecordWithInfoResult, err
}

/*
The AddRecordWithInfoWithReturnedRecordID method adds a record into the Senzing repository and returns information on the affected entities and the record identifier.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - jsonData: A JSON document containing the record to be added to the Senzing repository.
  - loadID: An identifier used to distinguish different load batches/sessions. An empty string is acceptable.
  - flags: Flags used to control information returned.

Output
  - A JSON document containing the AFFECTED_ENTITIES, INTERESTING_ENTITIES, and RECORD_ID.
    Example: `{"DATA_SOURCE":"TEST","RECORD_ID":"2D4DABB3FAEAFBD452E9487D06FABC22DC69C846","AFFECTED_ENTITIES":[{"ENTITY_ID":1}],"INTERESTING_ENTITIES":{"ENTITIES":[]}}`
  - The record identifier.
    Example: `2D4DABB3FAEAFBD452E9487D06FABC22DC69C846`
*/
func (client *G2engine) AddRecordWithInfoWithReturnedRecordID(ctx context.Context, dataSourceCode string, jsonData string, loadID string, flags int64) (string, string, error) {
	if client.isTrace {
		client.traceEntry(5, dataSourceCode, jsonData, loadID, flags)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"dataSourceCode": dataSourceCode,
				"recordID":       client.AddRecordWithInfoWithReturnedRecordIDResultRecordID,
				"loadID":         loadID,
			}
			client.notify(ctx, 8003, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(6, dataSourceCode, jsonData, loadID, flags, client.AddRecordWithInfoWithReturnedRecordIDResultGetWithInfo, client.AddRecordWithInfoWithReturnedRecordIDResultRecordID, err, time.Since(entryTime))
	}
	return client.AddRecordWithInfoWithReturnedRecordIDResultGetWithInfo, client.AddRecordWithInfoWithReturnedRecordIDResultRecordID, err
}

/*
The AddRecordWithReturnedRecordID method adds a record into the Senzing repository and returns the record identifier.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - jsonData: A JSON document containing the record to be added to the Senzing repository.
  - loadID: An identifier used to distinguish different load batches/sessions. An empty string is acceptable.

Output
  - The record identifier.
    Example: `2D4DABB3FAEAFBD452E9487D06FABC22DC69C846`
*/
func (client *G2engine) AddRecordWithReturnedRecordID(ctx context.Context, dataSourceCode string, jsonData string, loadID string) (string, error) {
	if client.isTrace {
		client.traceEntry(7, dataSourceCode, jsonData, loadID)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"dataSourceCode": dataSourceCode,
				"recordID":       client.AddRecordWithReturnedRecordIDResult,
				"loadID":         loadID,
			}
			client.notify(ctx, 8004, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(8, dataSourceCode, jsonData, loadID, client.AddRecordWithReturnedRecordIDResult, err, time.Since(entryTime))
	}
	return client.AddRecordWithReturnedRecordIDResult, err
}

/*
The CheckRecord method FIXME:.

Input
  - ctx: A context to control lifecycle.
  - record: A JSON document with the attribute data for the record to check with the "DATA_SOURCE" field.
  - recordQueryList: A JSON document with the datasource codes and recordID's of the records to check against.

Output

  - A JSON document that FIXME:
    See the example output.
*/
func (client *G2engine) CheckRecord(ctx context.Context, record string, recordQueryList string) (string, error) {
	if client.isTrace {
		client.traceEntry(9, record, recordQueryList)
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
		defer client.traceExit(10, record, recordQueryList, client.CheckRecordResult, err, time.Since(entryTime))
	}
	return client.CheckRecordResult, err
}

/*
The CloseExport method closes the exported document created by ExportJSONEntityReport().
It is part of the ExportJSONEntityReport(), FetchNext(), CloseExport()
lifecycle of a list of sized entities.

Input
  - ctx: A context to control lifecycle.
  - responseHandle: A handle created by ExportJSONEntityReport() or ExportCSVEntityReport().
*/
func (client *G2engine) CloseExport(ctx context.Context, responseHandle uintptr) error {
	if client.isTrace {
		client.traceEntry(13, responseHandle)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{}
			client.notify(ctx, 8006, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(14, responseHandle, err, time.Since(entryTime))
	}
	return err
}

/*
The CountRedoRecords method returns the number of records in need of redo-ing.

Input
  - ctx: A context to control lifecycle.

Output
  - The number of redo records in Senzing's redo queue.
*/
func (client *G2engine) CountRedoRecords(ctx context.Context) (int64, error) {
	if client.isTrace {
		client.traceEntry(15)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{}
			client.notify(ctx, 8007, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(16, client.CountRedoRecordsResult, err, time.Since(entryTime))
	}
	return client.CountRedoRecordsResult, err
}

/*
The DeleteRecord method deletes a record from the Senzing repository.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - recordID: The unique identifier within the records of the same data source.
  - loadID: An identifier used to distinguish different load batches/sessions. An empty string is acceptable.
    FIXME: How does the "loadID" affect what is deleted?
*/
func (client *G2engine) DeleteRecord(ctx context.Context, dataSourceCode string, recordID string, loadID string) error {
	if client.isTrace {
		client.traceEntry(17, dataSourceCode, recordID, loadID)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"dataSourceCode": dataSourceCode,
				"recordID":       recordID,
				"loadID":         loadID,
			}
			client.notify(ctx, 8008, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(18, dataSourceCode, recordID, loadID, err, time.Since(entryTime))
	}
	return err
}

/*
The DeleteRecordWithInfo method deletes a record from the Senzing repository and returns information on the affected entities.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - recordID: The unique identifier within the records of the same data source.
  - loadID: An identifier used to distinguish different load batches/sessions. An empty string is acceptable.
    FIXME: How does the "loadID" affect what is deleted?
  - flags: Flags used to control information returned.

Output
  - A JSON document.
    See the example output.
*/
func (client *G2engine) DeleteRecordWithInfo(ctx context.Context, dataSourceCode string, recordID string, loadID string, flags int64) (string, error) {
	if client.isTrace {
		client.traceEntry(19, dataSourceCode, recordID, loadID, flags)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"dataSourceCode": dataSourceCode,
				"recordID":       recordID,
				"loadID":         loadID,
			}
			client.notify(ctx, 8009, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(20, dataSourceCode, recordID, loadID, flags, client.DeleteRecordWithInfoResult, err, time.Since(entryTime))
	}
	return client.DeleteRecordWithInfoResult, err
}

/*
The Destroy method will destroy and perform cleanup for the Senzing G2 object.
It should be called after all other calls are complete.

Input
  - ctx: A context to control lifecycle.
*/
func (client *G2engine) Destroy(ctx context.Context) error {
	if client.isTrace {
		client.traceEntry(21)
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
		defer client.traceExit(22, err, time.Since(entryTime))
	}
	return err
}

/*
The ExportConfig method returns the Senzing engine configuration.

Input
  - ctx: A context to control lifecycle.

Output
  - A JSON document containing the current Senzing Engine configuration.
*/
func (client *G2engine) ExportConfig(ctx context.Context) (string, error) {
	if client.isTrace {
		client.traceEntry(25)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{}
			client.notify(ctx, 8011, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(26, client.ExportConfigResult, err, time.Since(entryTime))
	}
	return client.ExportConfigResult, err
}

/*
Similar to ExportConfig(), the ExportConfigAndConfigID method returns the Senzing engine configuration and it's identifier.

Input
  - ctx: A context to control lifecycle.

Output
  - A JSON document containing the current Senzing Engine configuration.
  - The unique identifier of the Senzing Engine configuration.
*/
func (client *G2engine) ExportConfigAndConfigID(ctx context.Context) (string, int64, error) {
	if client.isTrace {
		client.traceEntry(23)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"configID": strconv.FormatInt(client.ExportConfigAndConfigIDResultConfigID, 10),
			}
			client.notify(ctx, 8012, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(24, client.ExportConfigAndConfigIDResultConfig, client.ExportConfigAndConfigIDResultConfigID, err, time.Since(entryTime))
	}
	return client.ExportConfigAndConfigIDResultConfig, client.ExportConfigAndConfigIDResultConfigID, err
}

/*
The ExportCSVEntityReport method initializes a cursor over a document of exported entities.
It is part of the ExportCSVEntityReport(), FetchNext(), CloseExport()
lifecycle of a list of entities to export.

Input
  - ctx: A context to control lifecycle.
  - csvColumnList: A comma-separated list of column names for the CSV export.
  - flags: Flags used to control information returned.

Output
  - A handle that identifies the document to be scrolled through using FetchNext().
*/
func (client *G2engine) ExportCSVEntityReport(ctx context.Context, csvColumnList string, flags int64) (uintptr, error) {
	if client.isTrace {
		client.traceEntry(27, csvColumnList, flags)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{}
			client.notify(ctx, 8013, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(28, csvColumnList, flags, client.ExportCSVEntityReportResult, err, time.Since(entryTime))
	}
	return client.ExportCSVEntityReportResult, err
}

/*
The ExportJSONEntityReport method initializes a cursor over a document of exported entities.
It is part of the ExportJSONEntityReport(), FetchNext(), CloseExport()
lifecycle of a list of entities to export.

Input
  - ctx: A context to control lifecycle.
  - flags: Flags used to control information returned.

Output
  - A handle that identifies the document to be scrolled through using FetchNext().
*/
func (client *G2engine) ExportJSONEntityReport(ctx context.Context, flags int64) (uintptr, error) {
	if client.isTrace {
		client.traceEntry(29, flags)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{}
			client.notify(ctx, 8014, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(30, flags, client.ExportJSONEntityReportResult, err, time.Since(entryTime))
	}
	return client.ExportJSONEntityReportResult, err
}

/*
The FetchNext method is used to scroll through an exported document.
It is part of the ExportJSONEntityReport() or ExportCSVEntityReport(), FetchNext(), CloseExport()
lifecycle of a list of exported entities.

Input
  - ctx: A context to control lifecycle.
  - responseHandle: A handle created by ExportJSONEntityReport() or ExportCSVEntityReport().

Output
  - FIXME:
*/
func (client *G2engine) FetchNext(ctx context.Context, responseHandle uintptr) (string, error) {
	if client.isTrace {
		client.traceEntry(31, responseHandle)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{}
			client.notify(ctx, 8015, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(32, responseHandle, client.FetchNextResult, err, time.Since(entryTime))
	}
	return client.FetchNextResult, err
}

/*
The FindInterestingEntitiesByEntityID method FIXME:

Input
  - ctx: A context to control lifecycle.
  - entityID: The unique identifier of an entity.
  - flags: Flags used to control information returned.

Output
  - A JSON document.
    See the example output.
*/
func (client *G2engine) FindInterestingEntitiesByEntityID(ctx context.Context, entityID int64, flags int64) (string, error) {
	if client.isTrace {
		client.traceEntry(33, entityID, flags)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"entityID": strconv.FormatInt(entityID, 10),
			}
			client.notify(ctx, 8016, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(34, entityID, flags, client.FindInterestingEntitiesByEntityIDResult, err, time.Since(entryTime))
	}
	return client.FindInterestingEntitiesByEntityIDResult, err
}

/*
The FindInterestingEntitiesByRecordID method FIXME:

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - recordID: The unique identifier within the records of the same data source.
  - flags: Flags used to control information returned.

Output
  - A JSON document.
    See the example output.
*/
func (client *G2engine) FindInterestingEntitiesByRecordID(ctx context.Context, dataSourceCode string, recordID string, flags int64) (string, error) {
	if client.isTrace {
		client.traceEntry(35, dataSourceCode, recordID, flags)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"dataSourceCode": dataSourceCode,
				"recordID":       recordID,
			}
			client.notify(ctx, 8017, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(36, dataSourceCode, recordID, flags, client.FindInterestingEntitiesByRecordIDResult, err, time.Since(entryTime))
	}
	return client.FindInterestingEntitiesByRecordIDResult, err
}

/*
The FindNetworkByEntityID method finds all entities surrounding a requested set of entities.
This includes the requested entities, paths between them, and relations to other nearby entities.
To control output, use FindNetworkByEntityID_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - entityList: A JSON document listing entities.
    Example: `{"ENTITIES": [{"ENTITY_ID": 1}, {"ENTITY_ID": 2}, {"ENTITY_ID": 3}]}`
  - maxDegree: The maximum number of degrees in paths between search entities.
  - buildOutDegree: The number of degrees of relationships to show around each search entity.
  - maxEntities: The maximum number of entities to return in the discovered network.

Output
  - A JSON document.
    Example: `{"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":2,"ENTITIES":[1,2]}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1,"ENTITY_NAME":"SEAMAN","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":2,"FIRST_SEEN_DT":"2022-11-29 22:25:18.997","LAST_SEEN_DT":"2022-11-29 22:25:19.005"}],"LAST_SEEN_DT":"2022-11-29 22:25:19.005"},"RELATED_ENTITIES":[{"ENTITY_ID":2,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0}]},{"RESOLVED_ENTITY":{"ENTITY_ID":2,"ENTITY_NAME":"Smith","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-11-29 22:25:19.009","LAST_SEEN_DT":"2022-11-29 22:25:19.009"}],"LAST_SEEN_DT":"2022-11-29 22:25:19.009"},"RELATED_ENTITIES":[{"ENTITY_ID":1,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0}]}]}`
*/
func (client *G2engine) FindNetworkByEntityID(ctx context.Context, entityList string, maxDegree int, buildOutDegree int, maxEntities int) (string, error) {
	if client.isTrace {
		client.traceEntry(37, entityList, maxDegree, buildOutDegree, maxDegree)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"entityList": entityList,
			}
			client.notify(ctx, 8018, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(38, entityList, maxDegree, buildOutDegree, maxDegree, client.FindNetworkByEntityIDResult, err, time.Since(entryTime))
	}
	return client.FindNetworkByEntityIDResult, err
}

/*
The FindNetworkByEntityID_V2 method finds all entities surrounding a requested set of entities.
This includes the requested entities, paths between them, and relations to other nearby entities.
It extends FindNetworkByEntityID() by adding output control flags.

Input
  - ctx: A context to control lifecycle.
  - entityList: A JSON document listing entities.
    Example: `{"ENTITIES": [{"ENTITY_ID": 1}, {"ENTITY_ID": 2}, {"ENTITY_ID": 3}]}`
  - maxDegree: The maximum number of degrees in paths between search entities.
  - buildOutDegree: The number of degrees of relationships to show around each search entity.
  - maxEntities: The maximum number of entities to return in the discovered network.
  - flags: Flags used to control information returned.

Output
  - A JSON document.
    See the example output.
*/
func (client *G2engine) FindNetworkByEntityID_V2(ctx context.Context, entityList string, maxDegree int, buildOutDegree int, maxEntities int, flags int64) (string, error) {
	if client.isTrace {
		client.traceEntry(39, entityList, maxDegree, buildOutDegree, maxDegree, flags)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"entityList": entityList,
			}
			client.notify(ctx, 8019, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(40, entityList, maxDegree, buildOutDegree, maxDegree, flags, client.FindNetworkByEntityID_V2Result, err, time.Since(entryTime))
	}
	return client.FindNetworkByEntityID_V2Result, err
}

/*
The FindNetworkByRecordID method finds all entities surrounding a requested set of entities identified by record identifiers.
This includes the requested entities, paths between them, and relations to other nearby entities.
To control output, use FindNetworkByRecordID_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - entityList: A JSON document listing entities.
    Example: `{"ENTITIES": [{"ENTITY_ID": 1}, {"ENTITY_ID": 2}, {"ENTITY_ID": 3}]}`
  - maxDegree: The maximum number of degrees in paths between search entities.
  - buildOutDegree: The number of degrees of relationships to show around each search entity.
  - maxEntities: The maximum number of entities to return in the discovered network.

Output
  - A JSON document.
    Example: `{"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":2,"ENTITIES":[1,2]}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1,"ENTITY_NAME":"JOHNSON","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":2,"FIRST_SEEN_DT":"2022-12-06 14:40:34.285","LAST_SEEN_DT":"2022-12-06 14:40:34.420"}],"LAST_SEEN_DT":"2022-12-06 14:40:34.420"},"RELATED_ENTITIES":[{"ENTITY_ID":2,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0}]},{"RESOLVED_ENTITY":{"ENTITY_ID":2,"ENTITY_NAME":"OCEANGUY","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-12-06 14:40:34.359","LAST_SEEN_DT":"2022-12-06 14:40:34.359"}],"LAST_SEEN_DT":"2022-12-06 14:40:34.359"},"RELATED_ENTITIES":[{"ENTITY_ID":1,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+ADDRESS+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0}]},{"RESOLVED_ENTITY":{"ENTITY_ID":3,"ENTITY_NAME":"Smith","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-12-06 14:40:34.424","LAST_SEEN_DT":"2022-12-06 14:40:34.424"}],"LAST_SEEN_DT":"2022-12-06 14:40:34.424"},"RELATED_ENTITIES":[{"ENTITY_ID":1,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0},{"ENTITY_ID":2,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+ADDRESS+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0}]}]}`
*/
func (client *G2engine) FindNetworkByRecordID(ctx context.Context, recordList string, maxDegree int, buildOutDegree int, maxEntities int) (string, error) {
	if client.isTrace {
		client.traceEntry(41, recordList, maxDegree, buildOutDegree, maxDegree)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"recordList": recordList,
			}
			client.notify(ctx, 8020, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(42, recordList, maxDegree, buildOutDegree, maxDegree, client.FindNetworkByRecordIDResult, err, time.Since(entryTime))
	}
	return client.FindNetworkByRecordIDResult, err
}

/*
The FindNetworkByRecordID_V2 method finds all entities surrounding a requested set of entities identified by record identifiers.
This includes the requested entities, paths between them, and relations to other nearby entities.
It extends FindNetworkByRecordID() by adding output control flags.

Input
  - ctx: A context to control lifecycle.
  - entityList: A JSON document listing entities.
    Example: `{"ENTITIES": [{"ENTITY_ID": 1}, {"ENTITY_ID": 2}, {"ENTITY_ID": 3}]}`
  - maxDegree: The maximum number of degrees in paths between search entities.
  - buildOutDegree: The number of degrees of relationships to show around each search entity.
  - maxEntities: The maximum number of entities to return in the discovered network.
  - flags: Flags used to control information returned.

Output
  - A JSON document.
    See the example output.
*/
func (client *G2engine) FindNetworkByRecordID_V2(ctx context.Context, recordList string, maxDegree int, buildOutDegree int, maxEntities int, flags int64) (string, error) {
	if client.isTrace {
		client.traceEntry(43, recordList, maxDegree, buildOutDegree, maxDegree, flags)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"recordList": recordList,
			}
			client.notify(ctx, 8021, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(44, recordList, maxDegree, buildOutDegree, maxDegree, flags, client.FindNetworkByRecordID_V2Result, err, time.Since(entryTime))
	}
	return client.FindNetworkByRecordID_V2Result, err
}

/*
The FindPathByEntityID method finds single relationship paths between two entities.
Paths are found using known relationships with other entities.
To control output, use FindPathByEntityID_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - entityID1: The entity ID for the starting entity of the search path.
  - entityID2: The entity ID for the ending entity of the search path.
  - maxDegree: The maximum number of degrees in paths between search entities.

Output
  - A JSON document.
    Example: `{"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":2,"ENTITIES":[1,2]}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1,"ENTITY_NAME":"JOHNSON","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":2,"FIRST_SEEN_DT":"2022-12-06 14:43:49.024","LAST_SEEN_DT":"2022-12-06 14:43:49.164"}],"LAST_SEEN_DT":"2022-12-06 14:43:49.164"},"RELATED_ENTITIES":[{"ENTITY_ID":2,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0}]},{"RESOLVED_ENTITY":{"ENTITY_ID":2,"ENTITY_NAME":"OCEANGUY","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-12-06 14:43:49.104","LAST_SEEN_DT":"2022-12-06 14:43:49.104"}],"LAST_SEEN_DT":"2022-12-06 14:43:49.104"},"RELATED_ENTITIES":[{"ENTITY_ID":1,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+ADDRESS+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0}]}]}`
*/
func (client *G2engine) FindPathByEntityID(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int) (string, error) {
	if client.isTrace {
		client.traceEntry(45, entityID1, entityID2, maxDegree)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"entityID1": strconv.FormatInt(entityID1, 10),
				"entityID2": strconv.FormatInt(entityID2, 10),
			}
			client.notify(ctx, 8022, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(46, entityID1, entityID2, maxDegree, client.FindPathByEntityIDResult, err, time.Since(entryTime))
	}
	return client.FindPathByEntityIDResult, err
}

/*
The FindPathByEntityID_V2 method finds single relationship paths between two entities.
Paths are found using known relationships with other entities.
It extends FindPathByEntityID() by adding output control flags.

Input
  - ctx: A context to control lifecycle.
  - entityID1: The entity ID for the starting entity of the search path.
  - entityID2: The entity ID for the ending entity of the search path.
  - maxDegree: The maximum number of degrees in paths between search entities.
  - flags: Flags used to control information returned.

Output
  - A JSON document.
    See the example output.
*/
func (client *G2engine) FindPathByEntityID_V2(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int, flags int64) (string, error) {
	if client.isTrace {
		client.traceEntry(47, entityID1, entityID2, maxDegree, flags)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"entityID1": strconv.FormatInt(entityID1, 10),
				"entityID2": strconv.FormatInt(entityID2, 10),
			}
			client.notify(ctx, 8023, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(48, entityID1, entityID2, maxDegree, flags, client.FindPathByEntityID_V2Result, err, time.Since(entryTime))
	}
	return client.FindPathByEntityID_V2Result, err
}

/*
The FindPathByRecordID method finds single relationship paths between two entities.
The entities are identified by starting and ending records.
Paths are found using known relationships with other entities.
To control output, use FindPathByRecordID_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode1: Identifies the provenance of the record for the starting entity of the search path.
  - recordID1: The unique identifier within the records of the same data source for the starting entity of the search path.
  - dataSourceCode2: Identifies the provenance of the record for the ending entity of the search path.
  - recordID2: The unique identifier within the records of the same data source for the ending entity of the search path.
  - maxDegree: The maximum number of degrees in paths between search entities.

Output

  - A JSON document.
    Example: `{"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":2,"ENTITIES":[1,2]}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1,"ENTITY_NAME":"JOHNSON","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":2,"FIRST_SEEN_DT":"2022-12-06 14:48:19.522","LAST_SEEN_DT":"2022-12-06 14:48:19.667"}],"LAST_SEEN_DT":"2022-12-06 14:48:19.667"},"RELATED_ENTITIES":[{"ENTITY_ID":2,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0}]},{"RESOLVED_ENTITY":{"ENTITY_ID":2,"ENTITY_NAME":"OCEANGUY","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-12-06 14:48:19.593","LAST_SEEN_DT":"2022-12-06 14:48:19.593"}],"LAST_SEEN_DT":"2022-12-06 14:48:19.593"},"RELATED_ENTITIES":[{"ENTITY_ID":1,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+ADDRESS+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0}]}]}`
*/
func (client *G2engine) FindPathByRecordID(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int) (string, error) {
	if client.isTrace {
		client.traceEntry(49, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"dataSourceCode1": dataSourceCode1,
				"recordID1":       recordID1,
				"dataSourceCode2": dataSourceCode2,
				"recordID2":       recordID2,
			}
			client.notify(ctx, 8024, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(50, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, client.FindPathByRecordIDResult, err, time.Since(entryTime))
	}
	return client.FindPathByRecordIDResult, err
}

/*
The FindPathByRecordID_V2 method finds single relationship paths between two entities.
The entities are identified by starting and ending records.
Paths are found using known relationships with other entities.
It extends FindPathByRecordID() by adding output control flags.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode1: Identifies the provenance of the record for the starting entity of the search path.
  - recordID1: The unique identifier within the records of the same data source for the starting entity of the search path.
  - dataSourceCode2: Identifies the provenance of the record for the ending entity of the search path.
  - recordID2: The unique identifier within the records of the same data source for the ending entity of the search path.
  - maxDegree: The maximum number of degrees in paths between search entities.
  - flags: Flags used to control information returned.

Output
  - A JSON document.
    See the example output.
*/
func (client *G2engine) FindPathByRecordID_V2(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int, flags int64) (string, error) {
	if client.isTrace {
		client.traceEntry(51, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, flags)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"dataSourceCode1": dataSourceCode1,
				"recordID1":       recordID1,
				"dataSourceCode2": dataSourceCode2,
				"recordID2":       recordID2,
			}
			client.notify(ctx, 8025, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(52, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, flags, client.FindPathByRecordID_V2Result, err, time.Since(entryTime))
	}
	return client.FindPathByRecordID_V2Result, err
}

/*
The FindPathExcludingByEntityID method finds single relationship paths between two entities.
Paths are found using known relationships with other entities.
In addition, it will find paths that exclude certain entities from being on the path.
To control output, use FindPathExcludingByEntityID_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - entityID1: The entity ID for the starting entity of the search path.
  - entityID2: The entity ID for the ending entity of the search path.
  - maxDegree: The maximum number of degrees in paths between search entities.
  - excludedEntities: A JSON document listing entities that should be avoided on the path.

Output
  - A JSON document.
    Example: `{"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":2,"ENTITIES":[1,2]}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1,"ENTITY_NAME":"JOHNSON","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":2,"FIRST_SEEN_DT":"2022-12-06 14:50:49.222","LAST_SEEN_DT":"2022-12-06 14:50:49.356"}],"LAST_SEEN_DT":"2022-12-06 14:50:49.356"},"RELATED_ENTITIES":[{"ENTITY_ID":2,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0}]},{"RESOLVED_ENTITY":{"ENTITY_ID":2,"ENTITY_NAME":"OCEANGUY","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-12-06 14:50:49.295","LAST_SEEN_DT":"2022-12-06 14:50:49.295"}],"LAST_SEEN_DT":"2022-12-06 14:50:49.295"},"RELATED_ENTITIES":[{"ENTITY_ID":1,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+ADDRESS+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0}]}]}`
*/
func (client *G2engine) FindPathExcludingByEntityID(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int, excludedEntities string) (string, error) {
	if client.isTrace {
		client.traceEntry(53, entityID1, entityID2, maxDegree, excludedEntities)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"entityID1": strconv.FormatInt(entityID1, 10),
				"entityID2": strconv.FormatInt(entityID2, 10),
			}
			client.notify(ctx, 8026, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(54, entityID1, entityID2, maxDegree, excludedEntities, client.FindPathExcludingByEntityIDResult, err, time.Since(entryTime))
	}
	return client.FindPathExcludingByEntityIDResult, err
}

/*
The FindPathExcludingByEntityID_V2 method finds single relationship paths between two entities.
Paths are found using known relationships with other entities.
In addition, it will find paths that exclude certain entities from being on the path.
It extends FindPathExcludingByEntityID() by adding output control flags.

When excluding entities, the user may choose to either strictly exclude the entities,
or prefer to exclude the entities but still include them if no other path is found.
By default, entities will be strictly excluded.
A "preferred exclude" may be done by specifying the G2_FIND_PATH_PREFER_EXCLUDE control flag.

Input
  - ctx: A context to control lifecycle.
  - entityID1: The entity ID for the starting entity of the search path.
  - entityID2: The entity ID for the ending entity of the search path.
  - maxDegree: The maximum number of degrees in paths between search entities.
  - excludedEntities: A JSON document listing entities that should be avoided on the path.
  - flags: Flags used to control information returned.

Output
  - A JSON document.
    See the example output.
*/
func (client *G2engine) FindPathExcludingByEntityID_V2(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int, excludedEntities string, flags int64) (string, error) {
	if client.isTrace {
		client.traceEntry(55, entityID1, entityID2, maxDegree, excludedEntities, flags)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"entityID1": strconv.FormatInt(entityID1, 10),
				"entityID2": strconv.FormatInt(entityID2, 10),
			}
			client.notify(ctx, 8027, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(56, entityID1, entityID2, maxDegree, excludedEntities, flags, client.FindPathExcludingByEntityID_V2Result, err, time.Since(entryTime))
	}
	return client.FindPathExcludingByEntityID_V2Result, err
}

/*
The FindPathExcludingByRecordID method finds single relationship paths between two entities.
Paths are found using known relationships with other entities.
In addition, it will find paths that exclude certain entities from being on the path.
To control output, use FindPathExcludingByRecordID_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode1: Identifies the provenance of the record for the starting entity of the search path.
  - recordID1: The unique identifier within the records of the same data source for the starting entity of the search path.
  - dataSourceCode2: Identifies the provenance of the record for the ending entity of the search path.
  - recordID2: The unique identifier within the records of the same data source for the ending entity of the search path.
  - maxDegree: The maximum number of degrees in paths between search entities.
  - excludedRecords: A JSON document listing entities that should be avoided on the path.

Output
  - A JSON document.
    Example: `{"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":2,"ENTITIES":[1,2]}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1,"ENTITY_NAME":"JOHNSON","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":2,"FIRST_SEEN_DT":"2022-12-06 14:55:02.577","LAST_SEEN_DT":"2022-12-06 14:55:02.711"}],"LAST_SEEN_DT":"2022-12-06 14:55:02.711"},"RELATED_ENTITIES":[{"ENTITY_ID":2,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0}]},{"RESOLVED_ENTITY":{"ENTITY_ID":2,"ENTITY_NAME":"OCEANGUY","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-12-06 14:55:02.649","LAST_SEEN_DT":"2022-12-06 14:55:02.649"}],"LAST_SEEN_DT":"2022-12-06 14:55:02.649"},"RELATED_ENTITIES":[{"ENTITY_ID":1,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+ADDRESS+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0}]}]}`
*/
func (client *G2engine) FindPathExcludingByRecordID(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int, excludedRecords string) (string, error) {
	if client.isTrace {
		client.traceEntry(57, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, excludedRecords)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"dataSourceCode1": dataSourceCode1,
				"recordID1":       recordID1,
				"dataSourceCode2": dataSourceCode2,
				"recordID2":       recordID2,
			}
			client.notify(ctx, 8028, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(58, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, excludedRecords, client.FindPathExcludingByRecordIDResult, err, time.Since(entryTime))
	}
	return client.FindPathExcludingByRecordIDResult, err
}

/*
The FindPathExcludingByRecordID_V2 method finds single relationship paths between two entities.
Paths are found using known relationships with other entities.
In addition, it will find paths that exclude certain entities from being on the path.
It extends FindPathExcludingByRecordID() by adding output control flags.

When excluding entities, the user may choose to either strictly exclude the entities,
or prefer to exclude the entities but still include them if no other path is found.
By default, entities will be strictly excluded.
A "preferred exclude" may be done by specifying the G2_FIND_PATH_PREFER_EXCLUDE control flag.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode1: Identifies the provenance of the record for the starting entity of the search path.
  - recordID1: The unique identifier within the records of the same data source for the starting entity of the search path.
  - dataSourceCode2: Identifies the provenance of the record for the ending entity of the search path.
  - recordID2: The unique identifier within the records of the same data source for the ending entity of the search path.
  - maxDegree: The maximum number of degrees in paths between search entities.
  - excludedRecords: A JSON document listing entities that should be avoided on the path.
  - flags: Flags used to control information returned.

Output
  - A JSON document.
    See the example output.
*/
func (client *G2engine) FindPathExcludingByRecordID_V2(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int, excludedRecords string, flags int64) (string, error) {
	if client.isTrace {
		client.traceEntry(59, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, excludedRecords, flags)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"dataSourceCode1": dataSourceCode1,
				"recordID1":       recordID1,
				"dataSourceCode2": dataSourceCode2,
				"recordID2":       recordID2,
			}
			client.notify(ctx, 8029, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(60, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, excludedRecords, flags, client.FindPathExcludingByRecordID_V2Result, err, time.Since(entryTime))
	}
	return client.FindPathExcludingByRecordID_V2Result, err
}

/*
The FindPathIncludingSourceByEntityID method finds single relationship paths between two entities.
In addition, one of the enties along the path must include a specified data source.
Specific entities may also be excluded,
using the same methodology as the FindPathExcludingByEntityID() and FindPathExcludingByRecordID().
To control output, use FindPathIncludingSourceByEntityID_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - entityID1: The entity ID for the starting entity of the search path.
  - entityID2: The entity ID for the ending entity of the search path.
  - maxDegree: The maximum number of degrees in paths between search entities.
  - excludedEntities: A JSON document listing entities that should be avoided on the path.
  - requiredDsrcs: A JSON document listing data sources that should be included on the path. FIXME:

Output
  - A JSON document.
    Example: `{"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":2,"ENTITIES":[]}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1,"ENTITY_NAME":"JOHNSON","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":2,"FIRST_SEEN_DT":"2022-12-06 15:00:30.268","LAST_SEEN_DT":"2022-12-06 15:00:30.429"}],"LAST_SEEN_DT":"2022-12-06 15:00:30.429"},"RELATED_ENTITIES":[{"ENTITY_ID":2,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0}]},{"RESOLVED_ENTITY":{"ENTITY_ID":2,"ENTITY_NAME":"OCEANGUY","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-12-06 15:00:30.339","LAST_SEEN_DT":"2022-12-06 15:00:30.339"}],"LAST_SEEN_DT":"2022-12-06 15:00:30.339"},"RELATED_ENTITIES":[{"ENTITY_ID":1,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+ADDRESS+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0}]}]}`
*/
func (client *G2engine) FindPathIncludingSourceByEntityID(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int, excludedEntities string, requiredDsrcs string) (string, error) {
	if client.isTrace {
		client.traceEntry(61, entityID1, entityID2, maxDegree, excludedEntities, requiredDsrcs)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"entityID1": strconv.FormatInt(entityID1, 10),
				"entityID2": strconv.FormatInt(entityID2, 10),
			}
			client.notify(ctx, 8030, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(62, entityID1, entityID2, maxDegree, excludedEntities, requiredDsrcs, client.FindPathIncludingSourceByEntityIDResult, err, time.Since(entryTime))
	}
	return client.FindPathIncludingSourceByEntityIDResult, err
}

/*
The FindPathIncludingSourceByEntityID_V2 method finds single relationship paths between two entities.
In addition, one of the enties along the path must include a specified data source.
Specific entities may also be excluded,
using the same methodology as the FindPathExcludingByEntityID_V2() and FindPathExcludingByRecordID_V2().
It extends FindPathIncludingSourceByEntityID() by adding output control flags.

Input
  - ctx: A context to control lifecycle.
  - entityID1: The entity ID for the starting entity of the search path.
  - entityID2: The entity ID for the ending entity of the search path.
  - maxDegree: The maximum number of degrees in paths between search entities.
  - excludedEntities: A JSON document listing entities that should be avoided on the path.
  - requiredDsrcs: A JSON document listing data sources that should be included on the path. FIXME:
  - flags: Flags used to control information returned.

Output
  - A JSON document.
    See the example output.
*/
func (client *G2engine) FindPathIncludingSourceByEntityID_V2(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int, excludedEntities string, requiredDsrcs string, flags int64) (string, error) {
	if client.isTrace {
		client.traceEntry(63, entityID1, entityID2, maxDegree, excludedEntities, requiredDsrcs, flags)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"entityID1": strconv.FormatInt(entityID1, 10),
				"entityID2": strconv.FormatInt(entityID2, 10),
			}
			client.notify(ctx, 8031, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(64, entityID1, entityID2, maxDegree, excludedEntities, requiredDsrcs, flags, client.FindPathIncludingSourceByEntityID_V2Result, err, time.Since(entryTime))
	}
	return client.FindPathIncludingSourceByEntityID_V2Result, err
}

/*
The FindPathIncludingSourceByRecordID method finds single relationship paths between two entities.
In addition, one of the enties along the path must include a specified data source.
Specific entities may also be excluded,
using the same methodology as the FindPathExcludingByEntityID() and FindPathExcludingByRecordID().
To control output, use FindPathIncludingSourceByRecordID_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode1: Identifies the provenance of the record for the starting entity of the search path.
  - recordID1: The unique identifier within the records of the same data source for the starting entity of the search path.
  - dataSourceCode2: Identifies the provenance of the record for the ending entity of the search path.
  - recordID2: The unique identifier within the records of the same data source for the ending entity of the search path.
  - maxDegree: The maximum number of degrees in paths between search entities.
  - excludedRecords: A JSON document listing entities that should be avoided on the path.
  - requiredDsrcs: A JSON document listing data sources that should be included on the path. FIXME:

Output
  - A JSON document.
    Example: `{"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":2,"ENTITIES":[]}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1,"ENTITY_NAME":"JOHNSON","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":2,"FIRST_SEEN_DT":"2022-12-06 15:03:52.805","LAST_SEEN_DT":"2022-12-06 15:03:52.947"}],"LAST_SEEN_DT":"2022-12-06 15:03:52.947"},"RELATED_ENTITIES":[{"ENTITY_ID":2,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0}]},{"RESOLVED_ENTITY":{"ENTITY_ID":2,"ENTITY_NAME":"OCEANGUY","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-12-06 15:03:52.876","LAST_SEEN_DT":"2022-12-06 15:03:52.876"}],"LAST_SEEN_DT":"2022-12-06 15:03:52.876"},"RELATED_ENTITIES":[{"ENTITY_ID":1,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+ADDRESS+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0}]}]}`
*/
func (client *G2engine) FindPathIncludingSourceByRecordID(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int, excludedRecords string, requiredDsrcs string) (string, error) {
	if client.isTrace {
		client.traceEntry(65, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, excludedRecords, requiredDsrcs)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"dataSourceCode1": dataSourceCode1,
				"recordID1":       recordID1,
				"dataSourceCode2": dataSourceCode2,
				"recordID2":       recordID2,
			}
			client.notify(ctx, 8032, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(66, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, excludedRecords, requiredDsrcs, client.FindPathIncludingSourceByRecordIDResult, err, time.Since(entryTime))
	}
	return client.FindPathIncludingSourceByRecordIDResult, err
}

/*
The FindPathIncludingSourceByRecordID method finds single relationship paths between two entities.
In addition, one of the enties along the path must include a specified data source.
Specific entities may also be excluded,
using the same methodology as the FindPathExcludingByEntityID_V2() and FindPathExcludingByRecordID_V2().
It extends FindPathIncludingSourceByRecordID() by adding output control flags.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode1: Identifies the provenance of the record for the starting entity of the search path.
  - recordID1: The unique identifier within the records of the same data source for the starting entity of the search path.
  - dataSourceCode2: Identifies the provenance of the record for the ending entity of the search path.
  - recordID2: The unique identifier within the records of the same data source for the ending entity of the search path.
  - maxDegree: The maximum number of degrees in paths between search entities.
  - excludedRecords: A JSON document listing entities that should be avoided on the path.
  - requiredDsrcs: A JSON document listing data sources that should be included on the path. FIXME:
  - flags: Flags used to control information returned.

Output
  - A JSON document.
    See the example output.
*/
func (client *G2engine) FindPathIncludingSourceByRecordID_V2(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int, excludedRecords string, requiredDsrcs string, flags int64) (string, error) {
	if client.isTrace {
		client.traceEntry(67, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, excludedRecords, requiredDsrcs, flags)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"dataSourceCode1": dataSourceCode1,
				"recordID1":       recordID1,
				"dataSourceCode2": dataSourceCode2,
				"recordID2":       recordID2,
			}
			client.notify(ctx, 8033, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(68, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, excludedRecords, requiredDsrcs, flags, client.FindPathIncludingSourceByRecordID_V2Result, err, time.Since(entryTime))
	}
	return client.FindPathIncludingSourceByRecordID_V2Result, err
}

/*
The GetActiveConfigID method returns the identifier of the loaded Senzing engine configuration.

Input
  - ctx: A context to control lifecycle.

Output
  - The identifier of the active Senzing Engine configuration.
*/
func (client *G2engine) GetActiveConfigID(ctx context.Context) (int64, error) {
	if client.isTrace {
		client.traceEntry(69)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{}
			client.notify(ctx, 8034, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(70, client.GetActiveConfigIDResult, err, time.Since(entryTime))
	}
	return client.GetActiveConfigIDResult, err
}

/*
The GetEntityByEntityID method returns entity data based on the ID of a resolved identity.
To control output, use GetEntityByEntityID_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - entityID: The unique identifier of an entity.

Output

  - A JSON document.
    Example: `{"RESOLVED_ENTITY":{"ENTITY_ID":1,"ENTITY_NAME":"JOHNSON","FEATURES":{"ACCT_NUM":[{"FEAT_DESC":"5534202208773608","LIB_FEAT_ID":8,"USAGE_TYPE":"CC","FEAT_DESC_VALUES":[{"FEAT_DESC":"5534202208773608","LIB_FEAT_ID":8}]}],"ADDRESS":[{"FEAT_DESC":"772 Armstrong RD Delhi LA 71232","LIB_FEAT_ID":4,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772 Armstrong RD Delhi LA 71232","LIB_FEAT_ID":4}]}],"DOB":[{"FEAT_DESC":"4/8/1983","LIB_FEAT_ID":2,"FEAT_DESC_VALUES":[{"FEAT_DESC":"4/8/1983","LIB_FEAT_ID":2}]}],"GENDER":[{"FEAT_DESC":"F","LIB_FEAT_ID":3,"FEAT_DESC_VALUES":[{"FEAT_DESC":"F","LIB_FEAT_ID":3}]}],"LOGIN_ID":[{"FEAT_DESC":"flavorh","LIB_FEAT_ID":7,"FEAT_DESC_VALUES":[{"FEAT_DESC":"flavorh","LIB_FEAT_ID":7}]}],"NAME":[{"FEAT_DESC":"JOHNSON","LIB_FEAT_ID":1,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JOHNSON","LIB_FEAT_ID":1}]}],"PHONE":[{"FEAT_DESC":"225-671-0796","LIB_FEAT_ID":5,"FEAT_DESC_VALUES":[{"FEAT_DESC":"225-671-0796","LIB_FEAT_ID":5}]}],"SSN":[{"FEAT_DESC":"053-39-3251","LIB_FEAT_ID":6,"FEAT_DESC_VALUES":[{"FEAT_DESC":"053-39-3251","LIB_FEAT_ID":6}]}]},"RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":2,"FIRST_SEEN_DT":"2022-12-06 15:09:48.577","LAST_SEEN_DT":"2022-12-06 15:09:48.705"}],"LAST_SEEN_DT":"2022-12-06 15:09:48.705","RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":"111","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"C6063D4396612FBA7324DB0739273BA1FE815C43","ENTITY_DESC":"JOHNSON","MATCH_KEY":"","MATCH_LEVEL":0,"MATCH_LEVEL_CODE":"","ERRULE_CODE":"","LAST_SEEN_DT":"2022-12-06 15:09:48.577"},{"DATA_SOURCE":"TEST","RECORD_ID":"FCCE9793DAAD23159DBCCEB97FF2745B92CE7919","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"C6063D4396612FBA7324DB0739273BA1FE815C43","ENTITY_DESC":"JOHNSON","MATCH_KEY":"+EXACTLY_SAME","MATCH_LEVEL":0,"MATCH_LEVEL_CODE":"","ERRULE_CODE":"","LAST_SEEN_DT":"2022-12-06 15:09:48.705"}]},"RELATED_ENTITIES":[{"ENTITY_ID":2,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0,"ENTITY_NAME":"OCEANGUY","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-12-06 15:09:48.647","LAST_SEEN_DT":"2022-12-06 15:09:48.647"}],"LAST_SEEN_DT":"2022-12-06 15:09:48.647"},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0,"ENTITY_NAME":"Smith","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-12-06 15:09:48.709","LAST_SEEN_DT":"2022-12-06 15:09:48.709"}],"LAST_SEEN_DT":"2022-12-06 15:09:48.709"}]}`
*/
func (client *G2engine) GetEntityByEntityID(ctx context.Context, entityID int64) (string, error) {
	if client.isTrace {
		client.traceEntry(71, entityID)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"entityID": strconv.FormatInt(entityID, 10),
			}
			client.notify(ctx, 8035, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(72, entityID, client.GetEntityByEntityIDResult, err, time.Since(entryTime))
	}
	return client.GetEntityByEntityIDResult, err
}

/*
The GetEntityByEntityID_V2 method returns entity data based on the ID of a resolved identity.
It extends GetEntityByEntityID() by adding output control flags.

Input
  - ctx: A context to control lifecycle.
  - entityID: The unique identifier of an entity.
  - flags: Flags used to control information returned.

Output
  - A JSON document.
    See the example output.
*/
func (client *G2engine) GetEntityByEntityID_V2(ctx context.Context, entityID int64, flags int64) (string, error) {
	if client.isTrace {
		client.traceEntry(73, entityID, flags)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"entityID": strconv.FormatInt(entityID, 10),
			}
			client.notify(ctx, 8036, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(74, entityID, flags, client.GetEntityByEntityID_V2Result, err, time.Since(entryTime))
	}
	return client.GetEntityByEntityID_V2Result, err
}

/*
The GetEntityByRecordID method returns entity data based on the ID of a record which is a member of the entity.
To control output, use GetEntityByRecordID_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - recordID: The unique identifier within the records of the same data source.

Output
  - A JSON document.
    Example: `{"RESOLVED_ENTITY":{"ENTITY_ID":1,"ENTITY_NAME":"JOHNSON","FEATURES":{"ACCT_NUM":[{"FEAT_DESC":"5534202208773608","LIB_FEAT_ID":8,"USAGE_TYPE":"CC","FEAT_DESC_VALUES":[{"FEAT_DESC":"5534202208773608","LIB_FEAT_ID":8}]}],"ADDRESS":[{"FEAT_DESC":"772 Armstrong RD Delhi LA 71232","LIB_FEAT_ID":4,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772 Armstrong RD Delhi LA 71232","LIB_FEAT_ID":4}]}],"DOB":[{"FEAT_DESC":"4/8/1983","LIB_FEAT_ID":2,"FEAT_DESC_VALUES":[{"FEAT_DESC":"4/8/1983","LIB_FEAT_ID":2}]}],"GENDER":[{"FEAT_DESC":"F","LIB_FEAT_ID":3,"FEAT_DESC_VALUES":[{"FEAT_DESC":"F","LIB_FEAT_ID":3}]}],"LOGIN_ID":[{"FEAT_DESC":"flavorh","LIB_FEAT_ID":7,"FEAT_DESC_VALUES":[{"FEAT_DESC":"flavorh","LIB_FEAT_ID":7}]}],"NAME":[{"FEAT_DESC":"JOHNSON","LIB_FEAT_ID":1,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JOHNSON","LIB_FEAT_ID":1}]}],"PHONE":[{"FEAT_DESC":"225-671-0796","LIB_FEAT_ID":5,"FEAT_DESC_VALUES":[{"FEAT_DESC":"225-671-0796","LIB_FEAT_ID":5}]}],"SSN":[{"FEAT_DESC":"053-39-3251","LIB_FEAT_ID":6,"FEAT_DESC_VALUES":[{"FEAT_DESC":"053-39-3251","LIB_FEAT_ID":6}]}]},"RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":2,"FIRST_SEEN_DT":"2022-12-06 15:12:25.464","LAST_SEEN_DT":"2022-12-06 15:12:25.597"}],"LAST_SEEN_DT":"2022-12-06 15:12:25.597","RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":"111","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"C6063D4396612FBA7324DB0739273BA1FE815C43","ENTITY_DESC":"JOHNSON","MATCH_KEY":"","MATCH_LEVEL":0,"MATCH_LEVEL_CODE":"","ERRULE_CODE":"","LAST_SEEN_DT":"2022-12-06 15:12:25.464"},{"DATA_SOURCE":"TEST","RECORD_ID":"FCCE9793DAAD23159DBCCEB97FF2745B92CE7919","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"C6063D4396612FBA7324DB0739273BA1FE815C43","ENTITY_DESC":"JOHNSON","MATCH_KEY":"+EXACTLY_SAME","MATCH_LEVEL":0,"MATCH_LEVEL_CODE":"","ERRULE_CODE":"","LAST_SEEN_DT":"2022-12-06 15:12:25.597"}]},"RELATED_ENTITIES":[{"ENTITY_ID":2,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0,"ENTITY_NAME":"OCEANGUY","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-12-06 15:12:25.536","LAST_SEEN_DT":"2022-12-06 15:12:25.536"}],"LAST_SEEN_DT":"2022-12-06 15:12:25.536"},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0,"ENTITY_NAME":"Smith","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-12-06 15:12:25.603","LAST_SEEN_DT":"2022-12-06 15:12:25.603"}],"LAST_SEEN_DT":"2022-12-06 15:12:25.603"}]}`
*/
func (client *G2engine) GetEntityByRecordID(ctx context.Context, dataSourceCode string, recordID string) (string, error) {
	if client.isTrace {
		client.traceEntry(75, dataSourceCode, recordID)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"dataSourceCode": dataSourceCode,
				"recordID":       recordID,
			}
			client.notify(ctx, 8037, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(76, dataSourceCode, recordID, client.GetEntityByRecordIDResult, err, time.Since(entryTime))
	}
	return client.GetEntityByRecordIDResult, err
}

/*
The GetEntityByRecordID_V2 method returns entity data based on the ID of a record which is a member of the entity.
It extends GetEntityByRecordID() by adding output control flags.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - recordID: The unique identifier within the records of the same data source.
  - flags: Flags used to control information returned.

Output
  - A JSON document.
    See the example output.
*/
func (client *G2engine) GetEntityByRecordID_V2(ctx context.Context, dataSourceCode string, recordID string, flags int64) (string, error) {
	if client.isTrace {
		client.traceEntry(77, dataSourceCode, recordID, flags)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"dataSourceCode": dataSourceCode,
				"recordID":       recordID,
			}
			client.notify(ctx, 8038, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(78, dataSourceCode, recordID, flags, client.GetEntityByRecordID_V2Result, err, time.Since(entryTime))
	}
	return client.GetEntityByRecordID_V2Result, err
}

/*
The GetRecord method returns a JSON document of a single record from the Senzing repository.
To control output, use GetRecord_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - recordID: The unique identifier within the records of the same data source.

Output
  - A JSON document.
    See the example output.
*/
func (client *G2engine) GetRecord(ctx context.Context, dataSourceCode string, recordID string) (string, error) {
	if client.isTrace {
		client.traceEntry(83, dataSourceCode, recordID)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"dataSourceCode": dataSourceCode,
				"recordID":       recordID,
			}
			client.notify(ctx, 8039, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(84, dataSourceCode, recordID, client.GetRecordResult, err, time.Since(entryTime))
	}
	return client.GetRecordResult, err
}

/*
The GetRecord_V2 method returns a JSON document of a single record from the Senzing repository.
It extends GetRecord() by adding output control flags.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - recordID: The unique identifier within the records of the same data source.
  - flags: Flags used to control information returned.

Output
  - A JSON document.
    See the example output.
*/
func (client *G2engine) GetRecord_V2(ctx context.Context, dataSourceCode string, recordID string, flags int64) (string, error) {
	if client.isTrace {
		client.traceEntry(85, dataSourceCode, recordID, flags)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"dataSourceCode": dataSourceCode,
				"recordID":       recordID,
			}
			client.notify(ctx, 8040, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(86, dataSourceCode, recordID, flags, client.GetRecord_V2Result, err, time.Since(entryTime))
	}
	return client.GetRecord_V2Result, err
}

/*
The GetRedoRecord method returns the next internally queued maintenance record from the Senzing repository.
Usually, the ProcessRedoRecord() or ProcessRedoRecordWithInfo() method is called to process the maintenance record
retrieved by GetRedoRecord().

Input
  - ctx: A context to control lifecycle.

Output
  - A JSON document.
*/
func (client *G2engine) GetRedoRecord(ctx context.Context) (string, error) {
	if client.isTrace {
		client.traceEntry(87)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{}
			client.notify(ctx, 8041, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(88, client.GetRedoRecordResult, err, time.Since(entryTime))
	}
	return client.GetRedoRecordResult, err
}

/*
The GetRepositoryLastModifiedTime method retrieves the last modified time of the Senzing repository,
measured in the number of seconds between the last modified time and January 1, 1970 12:00am GMT (epoch time).

Input
  - ctx: A context to control lifecycle.

Output
  - A Unix Timestamp.
*/
func (client *G2engine) GetRepositoryLastModifiedTime(ctx context.Context) (int64, error) {
	if client.isTrace {
		client.traceEntry(89)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{}
			client.notify(ctx, 8042, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(90, client.GetRepositoryLastModifiedTimeResult, err, time.Since(entryTime))
	}
	return client.GetRepositoryLastModifiedTimeResult, err
}

/*
The GetSdkId method returns the identifier of this particular Software Development Kit (SDK).
It is handy when working with multiple implementations of the same G2engineInterface.
For this implementation, "mock" is returned.

Input
  - ctx: A context to control lifecycle.
*/
func (client *G2engine) GetSdkId(ctx context.Context) string {
	if client.isTrace {
		client.traceEntry(161)
	}
	entryTime := time.Now()
	var err error = nil
	if client.observers != nil {
		go func() {
			details := map[string]string{}
			client.notify(ctx, 8075, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(162, err, time.Since(entryTime))
	}
	return "mock"
}

/*
The GetVirtualEntityByRecordID method FIXME:
To control output, use GetVirtualEntityByRecordID_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - recordList: A JSON document.

Output
  - A JSON document.
    Example: `{"RESOLVED_ENTITY":{"ENTITY_ID":1,"ENTITY_NAME":"JOHNSON","FEATURES":{"ACCT_NUM":[{"FEAT_DESC":"5534202208773608","LIB_FEAT_ID":8,"USAGE_TYPE":"CC","FEAT_DESC_VALUES":[{"FEAT_DESC":"5534202208773608","LIB_FEAT_ID":8,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"ADDRESS":[{"FEAT_DESC":"772 Armstrong RD Delhi LA 71232","LIB_FEAT_ID":4,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772 Armstrong RD Delhi LA 71232","LIB_FEAT_ID":4,"USED_FOR_CAND":"N","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"772 Armstrong RD Delhi WI 53543","LIB_FEAT_ID":26,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772 Armstrong RD Delhi WI 53543","LIB_FEAT_ID":26,"USED_FOR_CAND":"N","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"ADDR_KEY":[{"FEAT_DESC":"772|ARMSTRNK||53543","LIB_FEAT_ID":37,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772|ARMSTRNK||53543","LIB_FEAT_ID":37,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"772|ARMSTRNK||71232","LIB_FEAT_ID":18,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772|ARMSTRNK||71232","LIB_FEAT_ID":18,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"772|ARMSTRNK||TL","LIB_FEAT_ID":17,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772|ARMSTRNK||TL","LIB_FEAT_ID":17,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"DOB":[{"FEAT_DESC":"4/8/1983","LIB_FEAT_ID":2,"FEAT_DESC_VALUES":[{"FEAT_DESC":"4/8/1983","LIB_FEAT_ID":2,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"6/9/1983","LIB_FEAT_ID":25,"FEAT_DESC_VALUES":[{"FEAT_DESC":"6/9/1983","LIB_FEAT_ID":25,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"GENDER":[{"FEAT_DESC":"F","LIB_FEAT_ID":3,"FEAT_DESC_VALUES":[{"FEAT_DESC":"F","LIB_FEAT_ID":3,"USED_FOR_CAND":"N","USED_FOR_SCORING":"Y","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"ID_KEY":[{"FEAT_DESC":"ACCT_NUM=5534202208773608","LIB_FEAT_ID":19,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ACCT_NUM=5534202208773608","LIB_FEAT_ID":19,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"SSN=053-39-3251","LIB_FEAT_ID":20,"FEAT_DESC_VALUES":[{"FEAT_DESC":"SSN=053-39-3251","LIB_FEAT_ID":20,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"SSN=153-33-5185","LIB_FEAT_ID":38,"FEAT_DESC_VALUES":[{"FEAT_DESC":"SSN=153-33-5185","LIB_FEAT_ID":38,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"LOGIN_ID":[{"FEAT_DESC":"flavorh","LIB_FEAT_ID":7,"FEAT_DESC_VALUES":[{"FEAT_DESC":"flavorh","LIB_FEAT_ID":7,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"flavorh2","LIB_FEAT_ID":28,"FEAT_DESC_VALUES":[{"FEAT_DESC":"flavorh2","LIB_FEAT_ID":28,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"NAME":[{"FEAT_DESC":"JOHNSON","LIB_FEAT_ID":1,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JOHNSON","LIB_FEAT_ID":1,"USED_FOR_CAND":"N","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"OCEANGUY","LIB_FEAT_ID":24,"FEAT_DESC_VALUES":[{"FEAT_DESC":"OCEANGUY","LIB_FEAT_ID":24,"USED_FOR_CAND":"N","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"NAME_KEY":[{"FEAT_DESC":"ASNK","LIB_FEAT_ID":29,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ASNK","LIB_FEAT_ID":29,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"ASNK|ADDRESS.CITY_STD=TL","LIB_FEAT_ID":34,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ASNK|ADDRESS.CITY_STD=TL","LIB_FEAT_ID":34,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"ASNK|DOB.MMDD_HASH=0906","LIB_FEAT_ID":32,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ASNK|DOB.MMDD_HASH=0906","LIB_FEAT_ID":32,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"ASNK|DOB.MMYY_HASH=0683","LIB_FEAT_ID":30,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ASNK|DOB.MMYY_HASH=0683","LIB_FEAT_ID":30,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"ASNK|DOB=80906","LIB_FEAT_ID":31,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ASNK|DOB=80906","LIB_FEAT_ID":31,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"ASNK|PHONE.PHONE_LAST_5=10796","LIB_FEAT_ID":33,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ASNK|PHONE.PHONE_LAST_5=10796","LIB_FEAT_ID":33,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"ASNK|POST=53543","LIB_FEAT_ID":36,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ASNK|POST=53543","LIB_FEAT_ID":36,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"ASNK|SSN=5185","LIB_FEAT_ID":35,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ASNK|SSN=5185","LIB_FEAT_ID":35,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN","LIB_FEAT_ID":11,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN","LIB_FEAT_ID":11,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|ADDRESS.CITY_STD=TL","LIB_FEAT_ID":12,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|ADDRESS.CITY_STD=TL","LIB_FEAT_ID":12,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|DOB.MMDD_HASH=0804","LIB_FEAT_ID":9,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|DOB.MMDD_HASH=0804","LIB_FEAT_ID":9,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|DOB.MMYY_HASH=0483","LIB_FEAT_ID":10,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|DOB.MMYY_HASH=0483","LIB_FEAT_ID":10,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|DOB=80804","LIB_FEAT_ID":13,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|DOB=80804","LIB_FEAT_ID":13,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|PHONE.PHONE_LAST_5=10796","LIB_FEAT_ID":15,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|PHONE.PHONE_LAST_5=10796","LIB_FEAT_ID":15,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|POST=71232","LIB_FEAT_ID":14,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|POST=71232","LIB_FEAT_ID":14,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|SSN=3251","LIB_FEAT_ID":16,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|SSN=3251","LIB_FEAT_ID":16,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"PHONE":[{"FEAT_DESC":"225-671-0796","LIB_FEAT_ID":5,"FEAT_DESC_VALUES":[{"FEAT_DESC":"225-671-0796","LIB_FEAT_ID":5,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"PHONE_KEY":[{"FEAT_DESC":"2256710796","LIB_FEAT_ID":21,"FEAT_DESC_VALUES":[{"FEAT_DESC":"2256710796","LIB_FEAT_ID":21,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"SEARCH_KEY":[{"FEAT_DESC":"LOGIN_ID:FLAVORH2|","LIB_FEAT_ID":40,"FEAT_DESC_VALUES":[{"FEAT_DESC":"LOGIN_ID:FLAVORH2|","LIB_FEAT_ID":40,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"LOGIN_ID:FLAVORH|","LIB_FEAT_ID":22,"FEAT_DESC_VALUES":[{"FEAT_DESC":"LOGIN_ID:FLAVORH|","LIB_FEAT_ID":22,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"SSN:3251|80804|","LIB_FEAT_ID":23,"FEAT_DESC_VALUES":[{"FEAT_DESC":"SSN:3251|80804|","LIB_FEAT_ID":23,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"SSN:5185|80906|","LIB_FEAT_ID":39,"FEAT_DESC_VALUES":[{"FEAT_DESC":"SSN:5185|80906|","LIB_FEAT_ID":39,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"SSN":[{"FEAT_DESC":"053-39-3251","LIB_FEAT_ID":6,"FEAT_DESC_VALUES":[{"FEAT_DESC":"053-39-3251","LIB_FEAT_ID":6,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"153-33-5185","LIB_FEAT_ID":27,"FEAT_DESC_VALUES":[{"FEAT_DESC":"153-33-5185","LIB_FEAT_ID":27,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}]},"RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":2,"FIRST_SEEN_DT":"2022-12-06 15:20:17.088","LAST_SEEN_DT":"2022-12-06 15:20:17.161"}],"LAST_SEEN_DT":"2022-12-06 15:20:17.161","RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":"111","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"C6063D4396612FBA7324DB0739273BA1FE815C43","ENTITY_DESC":"JOHNSON","LAST_SEEN_DT":"2022-12-06 15:20:17.088","FEATURES":[{"LIB_FEAT_ID":1},{"LIB_FEAT_ID":2},{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":4},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":6},{"LIB_FEAT_ID":7},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":9},{"LIB_FEAT_ID":10},{"LIB_FEAT_ID":11},{"LIB_FEAT_ID":12},{"LIB_FEAT_ID":13},{"LIB_FEAT_ID":14},{"LIB_FEAT_ID":15},{"LIB_FEAT_ID":16},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":18},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":20},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":22},{"LIB_FEAT_ID":23}]},{"DATA_SOURCE":"TEST","RECORD_ID":"222","ENTITY_TYPE":"TEST","INTERNAL_ID":2,"ENTITY_KEY":"740BA22D15CA88462A930AF8A7C904FF5E48226C","ENTITY_DESC":"OCEANGUY","LAST_SEEN_DT":"2022-12-06 15:20:17.161","FEATURES":[{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":24},{"LIB_FEAT_ID":25},{"LIB_FEAT_ID":26},{"LIB_FEAT_ID":27},{"LIB_FEAT_ID":28},{"LIB_FEAT_ID":29},{"LIB_FEAT_ID":30},{"LIB_FEAT_ID":31},{"LIB_FEAT_ID":32},{"LIB_FEAT_ID":33},{"LIB_FEAT_ID":34},{"LIB_FEAT_ID":35},{"LIB_FEAT_ID":36},{"LIB_FEAT_ID":37},{"LIB_FEAT_ID":38},{"LIB_FEAT_ID":39},{"LIB_FEAT_ID":40}]}]}}`
*/
func (client *G2engine) GetVirtualEntityByRecordID(ctx context.Context, recordList string) (string, error) {
	if client.isTrace {
		client.traceEntry(91, recordList)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"recordList": recordList,
			}
			client.notify(ctx, 8043, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(92, recordList, client.GetVirtualEntityByRecordIDResult, err, time.Since(entryTime))
	}
	return client.GetVirtualEntityByRecordIDResult, err
}

/*
The GetVirtualEntityByRecordID_V2 method FIXME:
It extends GetVirtualEntityByRecordID() by adding output control flags.

Input
  - ctx: A context to control lifecycle.
  - recordList: A JSON document.
    Example: `{"RECORDS": [{"DATA_SOURCE": "TEST","RECORD_ID": "111"},{"DATA_SOURCE": "TEST","RECORD_ID": "222"}]}`
  - flags: Flags used to control information returned.

Output
  - A JSON document.
    See the example output.
*/
func (client *G2engine) GetVirtualEntityByRecordID_V2(ctx context.Context, recordList string, flags int64) (string, error) {
	if client.isTrace {
		client.traceEntry(93, recordList, flags)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"recordList": recordList,
			}
			client.notify(ctx, 8044, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(94, recordList, flags, client.GetVirtualEntityByRecordID_V2Result, err, time.Since(entryTime))
	}
	return client.GetVirtualEntityByRecordID_V2Result, err
}

/*
The HowEntityByEntityID method FIXME:
To control output, use HowEntityByEntityID_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - entityID: The unique identifier of an entity.

Output
  - A JSON document.
    See the example output.
*/
func (client *G2engine) HowEntityByEntityID(ctx context.Context, entityID int64) (string, error) {
	if client.isTrace {
		client.traceEntry(95, entityID)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"entityID": strconv.FormatInt(entityID, 10),
			}
			client.notify(ctx, 8045, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(96, entityID, client.HowEntityByEntityIDResult, err, time.Since(entryTime))
	}
	return client.HowEntityByEntityIDResult, err
}

/*
The HowEntityByEntityID_V2 method FIXME:
It extends HowEntityByEntityID() by adding output control flags.

Input
  - ctx: A context to control lifecycle.
  - entityID: The unique identifier of an entity.
  - flags: Flags used to control information returned.

Output
  - A JSON document.
    See the example output.
*/
func (client *G2engine) HowEntityByEntityID_V2(ctx context.Context, entityID int64, flags int64) (string, error) {
	if client.isTrace {
		client.traceEntry(97, entityID, flags)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"entityID": strconv.FormatInt(entityID, 10),
			}
			client.notify(ctx, 8046, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(98, entityID, flags, client.HowEntityByEntityID_V2Result, err, time.Since(entryTime))
	}
	return client.HowEntityByEntityID_V2Result, err
}

/*
The Init method initializes the Senzing G2 object.
It must be called prior to any other calls.

Input
  - ctx: A context to control lifecycle.
  - moduleName: A name for the auditing node, to help identify it within system logs.
  - iniParams: A JSON string containing configuration parameters.
  - verboseLogging: A flag to enable deeper logging of the G2 processing. 0 for no Senzing logging; 1 for logging.
*/
func (client *G2engine) Init(ctx context.Context, moduleName string, iniParams string, verboseLogging int) error {
	if client.isTrace {
		client.traceEntry(99, moduleName, iniParams, verboseLogging)
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
			client.notify(ctx, 8047, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(100, moduleName, iniParams, verboseLogging, err, time.Since(entryTime))
	}
	return err
}

/*
The InitWithConfigID method initializes the Senzing G2 object with a non-default configuration ID.
It must be called prior to any other calls.

Input
  - ctx: A context to control lifecycle.
  - moduleName: A name for the auditing node, to help identify it within system logs.
  - iniParams: A JSON string containing configuration parameters.
  - initConfigID: The configuration ID used for the initialization.
  - verboseLogging: A flag to enable deeper logging of the G2 processing. 0 for no Senzing logging; 1 for logging.
*/
func (client *G2engine) InitWithConfigID(ctx context.Context, moduleName string, iniParams string, initConfigID int64, verboseLogging int) error {
	if client.isTrace {
		client.traceEntry(101, moduleName, iniParams, initConfigID, verboseLogging)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"iniParams":      iniParams,
				"initConfigID":   strconv.FormatInt(initConfigID, 10),
				"moduleName":     moduleName,
				"verboseLogging": strconv.Itoa(verboseLogging),
			}
			client.notify(ctx, 8048, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(102, moduleName, iniParams, initConfigID, verboseLogging, err, time.Since(entryTime))
	}
	return err
}

/*
The PrimeEngine method pre-initializes some of the heavier weight internal resources of the G2 engine.
The G2 Engine uses "lazy initialization".
PrimeEngine() forces initialization.

Input
  - ctx: A context to control lifecycle.
*/
func (client *G2engine) PrimeEngine(ctx context.Context) error {
	if client.isTrace {
		client.traceEntry(103)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{}
			client.notify(ctx, 8049, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(104, err, time.Since(entryTime))
	}
	return err
}

/*
The Process method FIXME:

Input
  - ctx: A context to control lifecycle.
  - record: A JSON document containing the record to be added to the Senzing repository.
*/
func (client *G2engine) Process(ctx context.Context, record string) error {
	if client.isTrace {
		client.traceEntry(105, record)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{}
			client.notify(ctx, 8050, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(106, record, err, time.Since(entryTime))
	}
	return err
}

/*
The ProcessRedoRecord method processes the next redo record and returns it.
Calling ProcessRedoRecord() has the potential to create more redo records in certain situations.

Input
  - ctx: A context to control lifecycle.

Output
  - A JSON document.
*/
func (client *G2engine) ProcessRedoRecord(ctx context.Context) (string, error) {
	if client.isTrace {
		client.traceEntry(107)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{}
			client.notify(ctx, 8051, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(108, client.ProcessRedoRecordResult, err, time.Since(entryTime))
	}
	return client.ProcessRedoRecordResult, err
}

/*
The ProcessRedoRecordWithInfo method processes the next redo record and returns it and affected entities.
Calling ProcessRedoRecordWithInfo() has the potential to create more redo records in certain situations.

Input
  - ctx: A context to control lifecycle.
  - flags: Flags used to control information returned.

Output
  - A JSON document with the record that was re-done.
  - A JSON document with affected entities.
*/
func (client *G2engine) ProcessRedoRecordWithInfo(ctx context.Context, flags int64) (string, string, error) {
	if client.isTrace {
		client.traceEntry(109, flags)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{}
			client.notify(ctx, 8052, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(110, flags, client.ProcessRedoRecordWithInfoResult, client.ProcessRedoRecordWithInfoResultWithInfo, err, time.Since(entryTime))
	}
	return client.ProcessRedoRecordWithInfoResult, client.ProcessRedoRecordWithInfoResultWithInfo, err
}

/*
The ProcessWithInfo method FIXME:

Input
  - ctx: A context to control lifecycle.
  - record: A JSON document containing the record to be added to the Senzing repository.
  - flags: Flags used to control information returned.

Output
  - A JSON document.
    See the example output.
*/
func (client *G2engine) ProcessWithInfo(ctx context.Context, record string, flags int64) (string, error) {
	if client.isTrace {
		client.traceEntry(111, record, flags)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{}
			client.notify(ctx, 8053, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(112, record, flags, client.ProcessWithInfoResult, err, time.Since(entryTime))
	}
	return client.ProcessWithInfoResult, err
}

/*
The ProcessWithResponse method FIXME:

Input
  - ctx: A context to control lifecycle.
  - record: A JSON document containing the record to be added to the Senzing repository.

Output
  - A JSON document.
    See the example output.
*/
func (client *G2engine) ProcessWithResponse(ctx context.Context, record string) (string, error) {
	if client.isTrace {
		client.traceEntry(113, record)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{}
			client.notify(ctx, 8054, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(114, record, client.ProcessWithResponseResult, err, time.Since(entryTime))
	}
	return client.ProcessWithResponseResult, err
}

/*
The ProcessWithResponseResize method FIXME:

Input
  - ctx: A context to control lifecycle.
  - record: A JSON document containing the record to be added to the Senzing repository.

Output
  - A JSON document.
    See the example output.
*/
func (client *G2engine) ProcessWithResponseResize(ctx context.Context, record string) (string, error) {
	if client.isTrace {
		client.traceEntry(115, record)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{}
			client.notify(ctx, 8055, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(116, record, client.ProcessWithResponseResizeResult, err, time.Since(entryTime))
	}
	return client.ProcessWithResponseResizeResult, err
}

/*
The PurgeRepository method removes every record in the Senzing repository.

Before calling purgeRepository() all other instances of the Senzing API
(whether in custom code, REST API, stream-loader, redoer, G2Loader, etc)
MUST be destroyed or shutdown.

Input
  - ctx: A context to control lifecycle.
*/
func (client *G2engine) PurgeRepository(ctx context.Context) error {
	if client.isTrace {
		client.traceEntry(117)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{}
			client.notify(ctx, 8056, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(118, err, time.Since(entryTime))
	}
	return err
}

/*
The ReevaluateEntity method FIXME:

Input
  - ctx: A context to control lifecycle.
  - entityID: The unique identifier of an entity.
  - flags: Flags used to control information returned.
*/
func (client *G2engine) ReevaluateEntity(ctx context.Context, entityID int64, flags int64) error {
	if client.isTrace {
		client.traceEntry(119, entityID, flags)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"entityID": strconv.FormatInt(entityID, 10),
			}
			client.notify(ctx, 8057, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(120, entityID, flags, err, time.Since(entryTime))
	}
	return err
}

/*
The ReevaluateEntityWithInfo method FIXME:

Input
  - ctx: A context to control lifecycle.
  - entityID: The unique identifier of an entity.
  - flags: Flags used to control information returned.

Output

  - A JSON document.
    See the example output.
*/
func (client *G2engine) ReevaluateEntityWithInfo(ctx context.Context, entityID int64, flags int64) (string, error) {
	if client.isTrace {
		client.traceEntry(121, entityID, flags)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"entityID": strconv.FormatInt(entityID, 10),
			}
			client.notify(ctx, 8058, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(122, entityID, flags, client.ReevaluateEntityWithInfoResult, err, time.Since(entryTime))
	}
	return client.ReevaluateEntityWithInfoResult, err
}

/*
The ReevaluateRecord method FIXME:

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - recordID: The unique identifier within the records of the same data source.
  - flags: Flags used to control information returned.
*/
func (client *G2engine) ReevaluateRecord(ctx context.Context, dataSourceCode string, recordID string, flags int64) error {
	if client.isTrace {
		client.traceEntry(123, dataSourceCode, recordID, flags)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"dataSourceCode": dataSourceCode,
				"recordID":       recordID,
			}
			client.notify(ctx, 8059, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(124, dataSourceCode, recordID, flags, err, time.Since(entryTime))
	}
	return err
}

/*
The ReevaluateRecordWithInfo method FIXME:

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - recordID: The unique identifier within the records of the same data source.
  - flags: Flags used to control information returned.

Output

  - A JSON document.
    See the example output.
*/
func (client *G2engine) ReevaluateRecordWithInfo(ctx context.Context, dataSourceCode string, recordID string, flags int64) (string, error) {
	if client.isTrace {
		client.traceEntry(125, dataSourceCode, recordID, flags)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"dataSourceCode": dataSourceCode,
				"recordID":       recordID,
			}
			client.notify(ctx, 8060, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(126, dataSourceCode, recordID, flags, client.ReevaluateRecordWithInfoResult, err, time.Since(entryTime))
	}
	return client.ReevaluateRecordWithInfoResult, err
}

/*
The RegisterObserver method adds the observer to the list of observers notified.

Input
  - ctx: A context to control lifecycle.
  - observer: The observer to be added.
*/
func (client *G2engine) RegisterObserver(ctx context.Context, observer observer.Observer) error {
	if client.isTrace {
		client.traceEntry(157, observer.GetObserverId(ctx))
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
			client.notify(ctx, 8076, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(158, observer.GetObserverId(ctx), err, time.Since(entryTime))
	}
	return err
}

/*
The Reinit method re-initializes the Senzing G2Engine object using a specified configuration identifier.

Input
  - ctx: A context to control lifecycle.
  - initConfigID: The configuration ID used for the initialization.
*/
func (client *G2engine) Reinit(ctx context.Context, initConfigID int64) error {
	if client.isTrace {
		client.traceEntry(127, initConfigID)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"initConfigID": strconv.FormatInt(initConfigID, 10),
			}
			client.notify(ctx, 8061, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(128, initConfigID, err, time.Since(entryTime))
	}
	return err
}

/*
The ReplaceRecord method updates/replaces a record in the Senzing repository.
If record doesn't exist, a new record is added to the data repository.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - recordID: The unique identifier within the records of the same data source.
  - jsonData: A JSON document containing the record to be added to the Senzing repository.
  - loadID: An identifier used to distinguish different load batches/sessions. An empty string is acceptable.
*/
func (client *G2engine) ReplaceRecord(ctx context.Context, dataSourceCode string, recordID string, jsonData string, loadID string) error {
	if client.isTrace {
		client.traceEntry(129, dataSourceCode, recordID, jsonData, loadID)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"dataSourceCode": dataSourceCode,
				"recordID":       recordID,
				"loadID":         loadID,
			}
			client.notify(ctx, 8062, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(130, dataSourceCode, recordID, jsonData, loadID, err, time.Since(entryTime))
	}
	return err
}

/*
The ReplaceRecordWithInfo method updates/replaces a record in the Senzing repository and returns information on the affected entities.
If record doesn't exist, a new record is added to the data repository.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - recordID: The unique identifier within the records of the same data source.
  - jsonData: A JSON document containing the record to be added to the Senzing repository.
  - loadID: An identifier used to distinguish different load batches/sessions. An empty string is acceptable.
  - flags: Flags used to control information returned.

Output
  - A JSON document.
    See the example output.
*/
func (client *G2engine) ReplaceRecordWithInfo(ctx context.Context, dataSourceCode string, recordID string, jsonData string, loadID string, flags int64) (string, error) {
	if client.isTrace {
		client.traceEntry(131, dataSourceCode, recordID, jsonData, loadID, flags)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"dataSourceCode": dataSourceCode,
				"recordID":       recordID,
				"loadID":         loadID,
			}
			client.notify(ctx, 8063, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(132, dataSourceCode, recordID, jsonData, loadID, flags, client.ReplaceRecordWithInfoResult, err, time.Since(entryTime))
	}
	return client.ReplaceRecordWithInfoResult, err
}

/*
The SearchByAttributes method retrieves entity data based on a user-specified set of entity attributes.
To control output, use SearchByAttributes_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - jsonData: A JSON document containing the record to be added to the Senzing repository.

Output
  - A JSON document.
    Example: `{"RESOLVED_ENTITIES":[{"MATCH_INFO":{"MATCH_LEVEL":1,"MATCH_LEVEL_CODE":"RESOLVED","MATCH_KEY":"+NAME+SSN","ERRULE_CODE":"SF1_PNAME_CSTAB","FEATURE_SCORES":{"NAME":[{"INBOUND_FEAT":"JOHNSON","CANDIDATE_FEAT":"JOHNSON","GNR_FN":100,"GNR_SN":100,"GNR_GN":70,"GENERATION_MATCH":-1,"GNR_ON":-1}],"SSN":[{"INBOUND_FEAT":"053-39-3251","CANDIDATE_FEAT":"053-39-3251","FULL_SCORE":100}]}},"ENTITY":{"RESOLVED_ENTITY":{"ENTITY_ID":1,"ENTITY_NAME":"JOHNSON","FEATURES":{"ACCT_NUM":[{"FEAT_DESC":"5534202208773608","LIB_FEAT_ID":8,"USAGE_TYPE":"CC","FEAT_DESC_VALUES":[{"FEAT_DESC":"5534202208773608","LIB_FEAT_ID":8}]}],"ADDRESS":[{"FEAT_DESC":"772 Armstrong RD Delhi LA 71232","LIB_FEAT_ID":4,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772 Armstrong RD Delhi LA 71232","LIB_FEAT_ID":4}]}],"DOB":[{"FEAT_DESC":"4/8/1983","LIB_FEAT_ID":2,"FEAT_DESC_VALUES":[{"FEAT_DESC":"4/8/1983","LIB_FEAT_ID":2}]},{"FEAT_DESC":"4/8/1985","LIB_FEAT_ID":100001,"FEAT_DESC_VALUES":[{"FEAT_DESC":"4/8/1985","LIB_FEAT_ID":100001}]}],"GENDER":[{"FEAT_DESC":"F","LIB_FEAT_ID":3,"FEAT_DESC_VALUES":[{"FEAT_DESC":"F","LIB_FEAT_ID":3}]}],"LOGIN_ID":[{"FEAT_DESC":"flavorh","LIB_FEAT_ID":7,"FEAT_DESC_VALUES":[{"FEAT_DESC":"flavorh","LIB_FEAT_ID":7}]}],"NAME":[{"FEAT_DESC":"JOHNSON","LIB_FEAT_ID":1,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JOHNSON","LIB_FEAT_ID":1}]}],"PHONE":[{"FEAT_DESC":"225-671-0796","LIB_FEAT_ID":5,"FEAT_DESC_VALUES":[{"FEAT_DESC":"225-671-0796","LIB_FEAT_ID":5}]}],"SSN":[{"FEAT_DESC":"053-39-3251","LIB_FEAT_ID":6,"FEAT_DESC_VALUES":[{"FEAT_DESC":"053-39-3251","LIB_FEAT_ID":6}]}]},"RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":6,"FIRST_SEEN_DT":"2022-12-06 15:38:06.175","LAST_SEEN_DT":"2022-12-06 15:38:06.957"}],"LAST_SEEN_DT":"2022-12-06 15:38:06.957"}}}]}`
*/
func (client *G2engine) SearchByAttributes(ctx context.Context, jsonData string) (string, error) {
	if client.isTrace {
		client.traceEntry(133, jsonData)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{}
			client.notify(ctx, 8064, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(134, jsonData, client.SearchByAttributesResult, err, time.Since(entryTime))
	}
	return client.SearchByAttributesResult, err
}

/*
The SearchByAttributes_V2 method retrieves entity data based on a user-specified set of entity attributes.
It extends SearchByAttributes() by adding output control flags.

Input
  - ctx: A context to control lifecycle.
  - jsonData: A JSON document containing the record to be added to the Senzing repository.
  - flags: Flags used to control information returned.

Output
  - A JSON document.
    See the example output.
*/
func (client *G2engine) SearchByAttributes_V2(ctx context.Context, jsonData string, flags int64) (string, error) {
	if client.isTrace {
		client.traceEntry(135, jsonData, flags)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{}
			client.notify(ctx, 8065, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(136, jsonData, flags, client.SearchByAttributes_V2Result, err, time.Since(entryTime))
	}
	return client.SearchByAttributes_V2Result, err
}

/*
The SetLogLevel method sets the level of logging.

Input
  - ctx: A context to control lifecycle.
  - logLevel: The desired log level. TRACE, DEBUG, INFO, WARN, ERROR, FATAL or PANIC.
*/
func (client *G2engine) SetLogLevel(ctx context.Context, logLevel logger.Level) error {
	if client.isTrace {
		client.traceEntry(137, logLevel)
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
			client.notify(ctx, 8077, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(138, logLevel, err, time.Since(entryTime))
	}
	return err
}

/*
The Stats method retrieves workload statistics for the current process.
These statistics will automatically reset after retrieval.

Input
  - ctx: A context to control lifecycle.

Output
  - A JSON document.
    Example: `{"workload":{"loadedRecords":5,"addedRecords":2,"deletedRecords":0,"reevaluations":0,"repairedEntities":0,"duration":56,"retries":0,"candidates":19,"actualAmbiguousTest":0,"cachedAmbiguousTest":0,"libFeatCacheHit":219,"libFeatCacheMiss":73,"unresolveTest":1,"abortedUnresolve":0,"gnrScorersUsed":1,"unresolveTriggers":{"normalResolve":0,"update":0,"relLink":0,"extensiveResolve":0,"ambiguousNoResolve":1,"ambiguousMultiResolve":0},"reresolveTriggers":{"abortRetry":0,"unresolveMovement":0,"multipleResolvableCandidates":0,"resolveNewFeatures":1,"newFeatureFTypes":[{"DOB":1}]},"reresolveSkipped":0,"filteredObsFeat":0,"expressedFeatureCalls":[{"EFCALL_ID":1,"EFUNC_CODE":"PHONE_HASHER","numCalls":1},{"EFCALL_ID":2,"EFUNC_CODE":"EXPRESS_ID","numCalls":1},{"EFCALL_ID":3,"EFUNC_CODE":"EXPRESS_ID","numCalls":1},{"EFCALL_ID":5,"EFUNC_CODE":"EXPRESS_BOM","numCalls":1},{"EFCALL_ID":7,"EFUNC_CODE":"NAME_HASHER","numCalls":4},{"EFCALL_ID":9,"EFUNC_CODE":"ADDR_HASHER","numCalls":1},{"EFCALL_ID":10,"EFUNC_CODE":"EXPRESS_BOM","numCalls":1},{"EFCALL_ID":14,"EFUNC_CODE":"EXPRESS_ID","numCalls":1},{"EFCALL_ID":16,"EFUNC_CODE":"EXPRESS_ID","numCalls":4}],"expressedFeaturesCreated":[{"ADDR_KEY":2},{"ID_KEY":7},{"NAME_KEY":14},{"PHONE_KEY":1},{"SEARCH_KEY":2}],"scoredPairs":[{"ACCT_NUM":16},{"ADDRESS":16},{"DOB":25},{"GENDER":16},{"LOGIN_ID":16},{"NAME":19},{"PHONE":16},{"SSN":19}],"cacheHit":[{"ADDRESS":12},{"DOB":18},{"NAME":13},{"PHONE":15}],"cacheMiss":[{"ADDRESS":4},{"DOB":7},{"NAME":6},{"PHONE":1}],"redoTriggers":[],"latchContention":[],"highContentionFeat":[],"highContentionResEnt":[],"genericDetect":[],"candidateBuilders":[{"ACCT_NUM":7},{"ADDR_KEY":7},{"DOB":7},{"ID_KEY":9},{"LOGIN_ID":7},{"NAME_KEY":9},{"PHONE":7},{"PHONE_KEY":7},{"SEARCH_KEY":7},{"SSN":9}],"suppressedCandidateBuilders":[],"suppressedScoredFeatureType":[],"reducedScoredFeatureType":[],"suppressedDisclosedRelationshipDomainCount":0,"CorruptEntityTestDiagnosis":{},"threadState":{"active":0,"idle":4,"sqlExecuting":0,"loader":0,"resolver":0,"scoring":0,"dataLatchContention":0,"obsEntContention":0,"resEntContention":0},"systemResources":{"initResources":[{"physicalCores":16},{"logicalCores":16},{"totalMemory":"62.6GB"},{"availableMemory":"49.5GB"}],"currResources":[{"availableMemory":"47.4GB"},{"activeThreads":0},{"workerThreads":4},{"systemLoad":[{"cpuUser":13.442277},{"cpuSystem":2.635741},{"cpuIdle":82.024246},{"cpuWait":1.634159},{"cpuSoftIrq":0.263574}]}]}}}`
*/
func (client *G2engine) Stats(ctx context.Context) (string, error) {
	if client.isTrace {
		client.traceEntry(139)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{}
			client.notify(ctx, 8066, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(140, client.StatsResult, err, time.Since(entryTime))
	}
	return client.StatsResult, err
}

/*
The UnregisterObserver method removes the observer to the list of observers notified.g2config

Input
  - ctx: A context to control lifecycle.
  - observer: The observer to be added.
*/
func (client *G2engine) UnregisterObserver(ctx context.Context, observer observer.Observer) error {
	if client.isTrace {
		client.traceEntry(159, observer.GetObserverId(ctx))
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
		client.notify(ctx, 8078, err, details)
	}
	err = client.observers.UnregisterObserver(ctx, observer)
	if !client.observers.HasObservers(ctx) {
		client.observers = nil
	}
	if client.isTrace {
		defer client.traceExit(160, observer.GetObserverId(ctx), err, time.Since(entryTime))
	}
	return err
}

/*
The WhyEntities method explains why records belong to their resolved entities.
WhyEntities() will compare the record data within an entity
against the rest of the entity data and show why they are connected.
This is calculated based on the features that record data represents.
To control output, use WhyEntities_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - entityID1: The entity ID for the starting entity of the search path.
  - entityID2: The entity ID for the ending entity of the search path.

Output
  - A JSON document.
    Example: `{"WHY_RESULTS":[{"ENTITY_ID":1,"ENTITY_ID_2":2,"MATCH_INFO":{"WHY_KEY":"+PHONE+ACCT_NUM-SSN","WHY_ERRULE_CODE":"SF1","MATCH_LEVEL_CODE":"POSSIBLY_RELATED","CANDIDATE_KEYS":{"ACCT_NUM":[{"FEAT_ID":8,"FEAT_DESC":"5534202208773608"}],"ADDR_KEY":[{"FEAT_ID":17,"FEAT_DESC":"772|ARMSTRNK||TL"}],"ID_KEY":[{"FEAT_ID":19,"FEAT_DESC":"ACCT_NUM=5534202208773608"}],"PHONE":[{"FEAT_ID":5,"FEAT_DESC":"225-671-0796"}],"PHONE_KEY":[{"FEAT_ID":21,"FEAT_DESC":"2256710796"}]},"DISCLOSED_RELATIONS":{},"FEATURE_SCORES":{"ACCT_NUM":[{"INBOUND_FEAT_ID":8,"INBOUND_FEAT":"5534202208773608","INBOUND_FEAT_USAGE_TYPE":"CC","CANDIDATE_FEAT_ID":8,"CANDIDATE_FEAT":"5534202208773608","CANDIDATE_FEAT_USAGE_TYPE":"CC","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"F1"}],"ADDRESS":[{"INBOUND_FEAT_ID":4,"INBOUND_FEAT":"772 Armstrong RD Delhi LA 71232","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":26,"CANDIDATE_FEAT":"772 Armstrong RD Delhi WI 53543","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":81,"SCORE_BUCKET":"LIKELY","SCORE_BEHAVIOR":"FF"}],"DOB":[{"INBOUND_FEAT_ID":100001,"INBOUND_FEAT":"4/8/1985","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":25,"CANDIDATE_FEAT":"6/9/1983","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":79,"SCORE_BUCKET":"NO_CHANCE","SCORE_BEHAVIOR":"FMES"},{"INBOUND_FEAT_ID":2,"INBOUND_FEAT":"4/8/1983","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":25,"CANDIDATE_FEAT":"6/9/1983","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":86,"SCORE_BUCKET":"PLAUSIBLE","SCORE_BEHAVIOR":"FMES"}],"GENDER":[{"INBOUND_FEAT_ID":3,"INBOUND_FEAT":"F","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":3,"CANDIDATE_FEAT":"F","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"FVME"}],"LOGIN_ID":[{"INBOUND_FEAT_ID":7,"INBOUND_FEAT":"flavorh","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":28,"CANDIDATE_FEAT":"flavorh2","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":0,"SCORE_BUCKET":"NO_CHANCE","SCORE_BEHAVIOR":"F1"}],"NAME":[{"INBOUND_FEAT_ID":1,"INBOUND_FEAT":"JOHNSON","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":24,"CANDIDATE_FEAT":"OCEANGUY","CANDIDATE_FEAT_USAGE_TYPE":"","GNR_FN":33,"GNR_SN":32,"GNR_GN":70,"GENERATION_MATCH":-1,"GNR_ON":-1,"SCORE_BUCKET":"NO_CHANCE","SCORE_BEHAVIOR":"NAME"}],"PHONE":[{"INBOUND_FEAT_ID":5,"INBOUND_FEAT":"225-671-0796","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":5,"CANDIDATE_FEAT":"225-671-0796","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"FF"}],"SSN":[{"INBOUND_FEAT_ID":6,"INBOUND_FEAT":"053-39-3251","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":27,"CANDIDATE_FEAT":"153-33-5185","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":0,"SCORE_BUCKET":"NO_CHANCE","SCORE_BEHAVIOR":"F1ES"}]}}}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1,"ENTITY_NAME":"JOHNSON","FEATURES":{"ACCT_NUM":[{"FEAT_DESC":"5534202208773608","LIB_FEAT_ID":8,"USAGE_TYPE":"CC","FEAT_DESC_VALUES":[{"FEAT_DESC":"5534202208773608","LIB_FEAT_ID":8,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"ADDRESS":[{"FEAT_DESC":"772 Armstrong RD Delhi LA 71232","LIB_FEAT_ID":4,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772 Armstrong RD Delhi LA 71232","LIB_FEAT_ID":4,"USED_FOR_CAND":"N","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"ADDR_KEY":[{"FEAT_DESC":"772|ARMSTRNK||71232","LIB_FEAT_ID":18,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772|ARMSTRNK||71232","LIB_FEAT_ID":18,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"772|ARMSTRNK||TL","LIB_FEAT_ID":17,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772|ARMSTRNK||TL","LIB_FEAT_ID":17,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"DOB":[{"FEAT_DESC":"4/8/1983","LIB_FEAT_ID":2,"FEAT_DESC_VALUES":[{"FEAT_DESC":"4/8/1983","LIB_FEAT_ID":2,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"4/8/1985","LIB_FEAT_ID":100001,"FEAT_DESC_VALUES":[{"FEAT_DESC":"4/8/1985","LIB_FEAT_ID":100001,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"GENDER":[{"FEAT_DESC":"F","LIB_FEAT_ID":3,"FEAT_DESC_VALUES":[{"FEAT_DESC":"F","LIB_FEAT_ID":3,"USED_FOR_CAND":"N","USED_FOR_SCORING":"Y","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"ID_KEY":[{"FEAT_DESC":"ACCT_NUM=5534202208773608","LIB_FEAT_ID":19,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ACCT_NUM=5534202208773608","LIB_FEAT_ID":19,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"SSN=053-39-3251","LIB_FEAT_ID":20,"FEAT_DESC_VALUES":[{"FEAT_DESC":"SSN=053-39-3251","LIB_FEAT_ID":20,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"LOGIN_ID":[{"FEAT_DESC":"flavorh","LIB_FEAT_ID":7,"FEAT_DESC_VALUES":[{"FEAT_DESC":"flavorh","LIB_FEAT_ID":7,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"NAME":[{"FEAT_DESC":"JOHNSON","LIB_FEAT_ID":1,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JOHNSON","LIB_FEAT_ID":1,"USED_FOR_CAND":"N","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"NAME_KEY":[{"FEAT_DESC":"JNSN","LIB_FEAT_ID":11,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN","LIB_FEAT_ID":11,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|ADDRESS.CITY_STD=TL","LIB_FEAT_ID":12,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|ADDRESS.CITY_STD=TL","LIB_FEAT_ID":12,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|DOB.MMDD_HASH=0804","LIB_FEAT_ID":9,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|DOB.MMDD_HASH=0804","LIB_FEAT_ID":9,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|DOB.MMYY_HASH=0483","LIB_FEAT_ID":10,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|DOB.MMYY_HASH=0483","LIB_FEAT_ID":10,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|DOB.MMYY_HASH=0485","LIB_FEAT_ID":100002,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|DOB.MMYY_HASH=0485","LIB_FEAT_ID":100002,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|DOB=80804","LIB_FEAT_ID":13,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|DOB=80804","LIB_FEAT_ID":13,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|PHONE.PHONE_LAST_5=10796","LIB_FEAT_ID":15,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|PHONE.PHONE_LAST_5=10796","LIB_FEAT_ID":15,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|POST=71232","LIB_FEAT_ID":14,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|POST=71232","LIB_FEAT_ID":14,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|SSN=3251","LIB_FEAT_ID":16,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|SSN=3251","LIB_FEAT_ID":16,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"PHONE":[{"FEAT_DESC":"225-671-0796","LIB_FEAT_ID":5,"FEAT_DESC_VALUES":[{"FEAT_DESC":"225-671-0796","LIB_FEAT_ID":5,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"PHONE_KEY":[{"FEAT_DESC":"2256710796","LIB_FEAT_ID":21,"FEAT_DESC_VALUES":[{"FEAT_DESC":"2256710796","LIB_FEAT_ID":21,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"SEARCH_KEY":[{"FEAT_DESC":"LOGIN_ID:FLAVORH|","LIB_FEAT_ID":22,"FEAT_DESC_VALUES":[{"FEAT_DESC":"LOGIN_ID:FLAVORH|","LIB_FEAT_ID":22,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"SSN:3251|80804|","LIB_FEAT_ID":23,"FEAT_DESC_VALUES":[{"FEAT_DESC":"SSN:3251|80804|","LIB_FEAT_ID":23,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"SSN":[{"FEAT_DESC":"053-39-3251","LIB_FEAT_ID":6,"FEAT_DESC_VALUES":[{"FEAT_DESC":"053-39-3251","LIB_FEAT_ID":6,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}]},"RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":6,"FIRST_SEEN_DT":"2022-12-06 15:58:57.129","LAST_SEEN_DT":"2022-12-06 15:58:57.906"}],"LAST_SEEN_DT":"2022-12-06 15:58:57.906","RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":"111","ENTITY_TYPE":"TEST","INTERNAL_ID":100001,"ENTITY_KEY":"A6C927986DF7329D1D2CDE0E8F34328AE640FB7E","ENTITY_DESC":"JOHNSON","MATCH_KEY":"","MATCH_LEVEL":0,"MATCH_LEVEL_CODE":"","ERRULE_CODE":"","LAST_SEEN_DT":"2022-12-06 15:58:57.906","FEATURES":[{"LIB_FEAT_ID":1},{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":4},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":6},{"LIB_FEAT_ID":7},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":9},{"LIB_FEAT_ID":11},{"LIB_FEAT_ID":12},{"LIB_FEAT_ID":13},{"LIB_FEAT_ID":14},{"LIB_FEAT_ID":15},{"LIB_FEAT_ID":16},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":18},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":20},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":22},{"LIB_FEAT_ID":23},{"LIB_FEAT_ID":100001},{"LIB_FEAT_ID":100002}]},{"DATA_SOURCE":"TEST","RECORD_ID":"444","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"C6063D4396612FBA7324DB0739273BA1FE815C43","ENTITY_DESC":"JOHNSON","MATCH_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","MATCH_LEVEL":1,"MATCH_LEVEL_CODE":"RESOLVED","ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","LAST_SEEN_DT":"2022-12-06 15:58:57.400","FEATURES":[{"LIB_FEAT_ID":1},{"LIB_FEAT_ID":2},{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":4},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":6},{"LIB_FEAT_ID":7},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":9},{"LIB_FEAT_ID":10},{"LIB_FEAT_ID":11},{"LIB_FEAT_ID":12},{"LIB_FEAT_ID":13},{"LIB_FEAT_ID":14},{"LIB_FEAT_ID":15},{"LIB_FEAT_ID":16},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":18},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":20},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":22},{"LIB_FEAT_ID":23}]},{"DATA_SOURCE":"TEST","RECORD_ID":"555","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"C6063D4396612FBA7324DB0739273BA1FE815C43","ENTITY_DESC":"JOHNSON","MATCH_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","MATCH_LEVEL":1,"MATCH_LEVEL_CODE":"RESOLVED","ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","LAST_SEEN_DT":"2022-12-06 15:58:57.404","FEATURES":[{"LIB_FEAT_ID":1},{"LIB_FEAT_ID":2},{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":4},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":6},{"LIB_FEAT_ID":7},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":9},{"LIB_FEAT_ID":10},{"LIB_FEAT_ID":11},{"LIB_FEAT_ID":12},{"LIB_FEAT_ID":13},{"LIB_FEAT_ID":14},{"LIB_FEAT_ID":15},{"LIB_FEAT_ID":16},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":18},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":20},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":22},{"LIB_FEAT_ID":23}]},{"DATA_SOURCE":"TEST","RECORD_ID":"666","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"C6063D4396612FBA7324DB0739273BA1FE815C43","ENTITY_DESC":"JOHNSON","MATCH_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","MATCH_LEVEL":1,"MATCH_LEVEL_CODE":"RESOLVED","ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","LAST_SEEN_DT":"2022-12-06 15:58:57.407","FEATURES":[{"LIB_FEAT_ID":1},{"LIB_FEAT_ID":2},{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":4},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":6},{"LIB_FEAT_ID":7},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":9},{"LIB_FEAT_ID":10},{"LIB_FEAT_ID":11},{"LIB_FEAT_ID":12},{"LIB_FEAT_ID":13},{"LIB_FEAT_ID":14},{"LIB_FEAT_ID":15},{"LIB_FEAT_ID":16},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":18},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":20},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":22},{"LIB_FEAT_ID":23}]},{"DATA_SOURCE":"TEST","RECORD_ID":"777","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"C6063D4396612FBA7324DB0739273BA1FE815C43","ENTITY_DESC":"JOHNSON","MATCH_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","MATCH_LEVEL":1,"MATCH_LEVEL_CODE":"RESOLVED","ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","LAST_SEEN_DT":"2022-12-06 15:58:57.410","FEATURES":[{"LIB_FEAT_ID":1},{"LIB_FEAT_ID":2},{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":4},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":6},{"LIB_FEAT_ID":7},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":9},{"LIB_FEAT_ID":10},{"LIB_FEAT_ID":11},{"LIB_FEAT_ID":12},{"LIB_FEAT_ID":13},{"LIB_FEAT_ID":14},{"LIB_FEAT_ID":15},{"LIB_FEAT_ID":16},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":18},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":20},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":22},{"LIB_FEAT_ID":23}]},{"DATA_SOURCE":"TEST","RECORD_ID":"FCCE9793DAAD23159DBCCEB97FF2745B92CE7919","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"C6063D4396612FBA7324DB0739273BA1FE815C43","ENTITY_DESC":"JOHNSON","MATCH_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","MATCH_LEVEL":1,"MATCH_LEVEL_CODE":"RESOLVED","ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","LAST_SEEN_DT":"2022-12-06 15:58:57.259","FEATURES":[{"LIB_FEAT_ID":1},{"LIB_FEAT_ID":2},{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":4},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":6},{"LIB_FEAT_ID":7},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":9},{"LIB_FEAT_ID":10},{"LIB_FEAT_ID":11},{"LIB_FEAT_ID":12},{"LIB_FEAT_ID":13},{"LIB_FEAT_ID":14},{"LIB_FEAT_ID":15},{"LIB_FEAT_ID":16},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":18},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":20},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":22},{"LIB_FEAT_ID":23}]}]},"RELATED_ENTITIES":[{"ENTITY_ID":2,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0,"ENTITY_NAME":"OCEANGUY","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-12-06 15:58:57.201","LAST_SEEN_DT":"2022-12-06 15:58:57.201"}],"LAST_SEEN_DT":"2022-12-06 15:58:57.201"},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0,"ENTITY_NAME":"Smith","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-12-06 15:58:57.263","LAST_SEEN_DT":"2022-12-06 15:58:57.263"}],"LAST_SEEN_DT":"2022-12-06 15:58:57.263"}]},{"RESOLVED_ENTITY":{"ENTITY_ID":2,"ENTITY_NAME":"OCEANGUY","FEATURES":{"ACCT_NUM":[{"FEAT_DESC":"5534202208773608","LIB_FEAT_ID":8,"USAGE_TYPE":"CC","FEAT_DESC_VALUES":[{"FEAT_DESC":"5534202208773608","LIB_FEAT_ID":8,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"ADDRESS":[{"FEAT_DESC":"772 Armstrong RD Delhi WI 53543","LIB_FEAT_ID":26,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772 Armstrong RD Delhi WI 53543","LIB_FEAT_ID":26,"USED_FOR_CAND":"N","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"ADDR_KEY":[{"FEAT_DESC":"772|ARMSTRNK||53543","LIB_FEAT_ID":37,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772|ARMSTRNK||53543","LIB_FEAT_ID":37,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"772|ARMSTRNK||TL","LIB_FEAT_ID":17,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772|ARMSTRNK||TL","LIB_FEAT_ID":17,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"DOB":[{"FEAT_DESC":"6/9/1983","LIB_FEAT_ID":25,"FEAT_DESC_VALUES":[{"FEAT_DESC":"6/9/1983","LIB_FEAT_ID":25,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"GENDER":[{"FEAT_DESC":"F","LIB_FEAT_ID":3,"FEAT_DESC_VALUES":[{"FEAT_DESC":"F","LIB_FEAT_ID":3,"USED_FOR_CAND":"N","USED_FOR_SCORING":"Y","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"ID_KEY":[{"FEAT_DESC":"ACCT_NUM=5534202208773608","LIB_FEAT_ID":19,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ACCT_NUM=5534202208773608","LIB_FEAT_ID":19,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"SSN=153-33-5185","LIB_FEAT_ID":38,"FEAT_DESC_VALUES":[{"FEAT_DESC":"SSN=153-33-5185","LIB_FEAT_ID":38,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"LOGIN_ID":[{"FEAT_DESC":"flavorh2","LIB_FEAT_ID":28,"FEAT_DESC_VALUES":[{"FEAT_DESC":"flavorh2","LIB_FEAT_ID":28,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"NAME":[{"FEAT_DESC":"OCEANGUY","LIB_FEAT_ID":24,"FEAT_DESC_VALUES":[{"FEAT_DESC":"OCEANGUY","LIB_FEAT_ID":24,"USED_FOR_CAND":"N","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"NAME_KEY":[{"FEAT_DESC":"ASNK","LIB_FEAT_ID":29,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ASNK","LIB_FEAT_ID":29,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"ASNK|ADDRESS.CITY_STD=TL","LIB_FEAT_ID":34,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ASNK|ADDRESS.CITY_STD=TL","LIB_FEAT_ID":34,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"ASNK|DOB.MMDD_HASH=0906","LIB_FEAT_ID":32,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ASNK|DOB.MMDD_HASH=0906","LIB_FEAT_ID":32,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"ASNK|DOB.MMYY_HASH=0683","LIB_FEAT_ID":30,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ASNK|DOB.MMYY_HASH=0683","LIB_FEAT_ID":30,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"ASNK|DOB=80906","LIB_FEAT_ID":31,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ASNK|DOB=80906","LIB_FEAT_ID":31,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"ASNK|PHONE.PHONE_LAST_5=10796","LIB_FEAT_ID":33,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ASNK|PHONE.PHONE_LAST_5=10796","LIB_FEAT_ID":33,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"ASNK|POST=53543","LIB_FEAT_ID":36,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ASNK|POST=53543","LIB_FEAT_ID":36,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"ASNK|SSN=5185","LIB_FEAT_ID":35,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ASNK|SSN=5185","LIB_FEAT_ID":35,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"PHONE":[{"FEAT_DESC":"225-671-0796","LIB_FEAT_ID":5,"FEAT_DESC_VALUES":[{"FEAT_DESC":"225-671-0796","LIB_FEAT_ID":5,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"PHONE_KEY":[{"FEAT_DESC":"2256710796","LIB_FEAT_ID":21,"FEAT_DESC_VALUES":[{"FEAT_DESC":"2256710796","LIB_FEAT_ID":21,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"SEARCH_KEY":[{"FEAT_DESC":"LOGIN_ID:FLAVORH2|","LIB_FEAT_ID":40,"FEAT_DESC_VALUES":[{"FEAT_DESC":"LOGIN_ID:FLAVORH2|","LIB_FEAT_ID":40,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"SSN:5185|80906|","LIB_FEAT_ID":39,"FEAT_DESC_VALUES":[{"FEAT_DESC":"SSN:5185|80906|","LIB_FEAT_ID":39,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"SSN":[{"FEAT_DESC":"153-33-5185","LIB_FEAT_ID":27,"FEAT_DESC_VALUES":[{"FEAT_DESC":"153-33-5185","LIB_FEAT_ID":27,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}]},"RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-12-06 15:58:57.201","LAST_SEEN_DT":"2022-12-06 15:58:57.201"}],"LAST_SEEN_DT":"2022-12-06 15:58:57.201","RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":"222","ENTITY_TYPE":"TEST","INTERNAL_ID":2,"ENTITY_KEY":"740BA22D15CA88462A930AF8A7C904FF5E48226C","ENTITY_DESC":"OCEANGUY","MATCH_KEY":"","MATCH_LEVEL":0,"MATCH_LEVEL_CODE":"","ERRULE_CODE":"","LAST_SEEN_DT":"2022-12-06 15:58:57.201","FEATURES":[{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":24},{"LIB_FEAT_ID":25},{"LIB_FEAT_ID":26},{"LIB_FEAT_ID":27},{"LIB_FEAT_ID":28},{"LIB_FEAT_ID":29},{"LIB_FEAT_ID":30},{"LIB_FEAT_ID":31},{"LIB_FEAT_ID":32},{"LIB_FEAT_ID":33},{"LIB_FEAT_ID":34},{"LIB_FEAT_ID":35},{"LIB_FEAT_ID":36},{"LIB_FEAT_ID":37},{"LIB_FEAT_ID":38},{"LIB_FEAT_ID":39},{"LIB_FEAT_ID":40}]}]},"RELATED_ENTITIES":[{"ENTITY_ID":1,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0,"ENTITY_NAME":"JOHNSON","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":6,"FIRST_SEEN_DT":"2022-12-06 15:58:57.129","LAST_SEEN_DT":"2022-12-06 15:58:57.906"}],"LAST_SEEN_DT":"2022-12-06 15:58:57.906"},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+ADDRESS+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0,"ENTITY_NAME":"Smith","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-12-06 15:58:57.263","LAST_SEEN_DT":"2022-12-06 15:58:57.263"}],"LAST_SEEN_DT":"2022-12-06 15:58:57.263"}]}]}`
*/
func (client *G2engine) WhyEntities(ctx context.Context, entityID1 int64, entityID2 int64) (string, error) {
	if client.isTrace {
		client.traceEntry(141, entityID1, entityID2)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"entityID1": strconv.FormatInt(entityID1, 10),
				"entityID2": strconv.FormatInt(entityID2, 10),
			}
			client.notify(ctx, 8067, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(142, entityID1, entityID2, client.WhyEntitiesResult, err, time.Since(entryTime))
	}
	return client.WhyEntitiesResult, err
}

/*
The WhyEntities_V2 method explains why records belong to their resolved entities.
WhyEntities_V2() will compare the record data within an entity
against the rest of the entity data and show why they are connected.
This is calculated based on the features that record data represents.
It extends WhyEntities() by adding output control flags.

Input
  - ctx: A context to control lifecycle.
  - entityID1: The entity ID for the starting entity of the search path.
  - entityID2: The entity ID for the ending entity of the search path.
  - flags: Flags used to control information returned.

Output
  - A JSON document.
    See the example output.
*/
func (client *G2engine) WhyEntities_V2(ctx context.Context, entityID1 int64, entityID2 int64, flags int64) (string, error) {
	if client.isTrace {
		client.traceEntry(143, entityID1, entityID2, flags)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"entityID1": strconv.FormatInt(entityID1, 10),
				"entityID2": strconv.FormatInt(entityID2, 10),
			}
			client.notify(ctx, 8068, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(144, entityID1, entityID2, flags, client.WhyEntities_V2Result, err, time.Since(entryTime))
	}
	return client.WhyEntities_V2Result, err
}

/*
The WhyEntityByEntityID method explains why records belong to their resolved entities.
To control output, use WhyEntityByEntityID_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - entityID: The unique identifier of an entity for the starting entity of the search path.

Output

  - A JSON document.
    Example: `{"WHY_RESULTS":[{"INTERNAL_ID":1,"ENTITY_ID":1,"FOCUS_RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":"444"},{"DATA_SOURCE":"TEST","RECORD_ID":"555"},{"DATA_SOURCE":"TEST","RECORD_ID":"666"},{"DATA_SOURCE":"TEST","RECORD_ID":"777"},{"DATA_SOURCE":"TEST","RECORD_ID":"FCCE9793DAAD23159DBCCEB97FF2745B92CE7919"}],"MATCH_INFO":{"WHY_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","WHY_ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","MATCH_LEVEL_CODE":"RESOLVED","CANDIDATE_KEYS":{"ACCT_NUM":[{"FEAT_ID":8,"FEAT_DESC":"5534202208773608"}],"ADDR_KEY":[{"FEAT_ID":17,"FEAT_DESC":"772|ARMSTRNK||TL"},{"FEAT_ID":18,"FEAT_DESC":"772|ARMSTRNK||71232"}],"ID_KEY":[{"FEAT_ID":19,"FEAT_DESC":"ACCT_NUM=5534202208773608"},{"FEAT_ID":20,"FEAT_DESC":"SSN=053-39-3251"}],"LOGIN_ID":[{"FEAT_ID":7,"FEAT_DESC":"flavorh"}],"NAME_KEY":[{"FEAT_ID":9,"FEAT_DESC":"JNSN|DOB.MMDD_HASH=0804"},{"FEAT_ID":11,"FEAT_DESC":"JNSN"},{"FEAT_ID":12,"FEAT_DESC":"JNSN|ADDRESS.CITY_STD=TL"},{"FEAT_ID":13,"FEAT_DESC":"JNSN|DOB=80804"},{"FEAT_ID":14,"FEAT_DESC":"JNSN|POST=71232"},{"FEAT_ID":15,"FEAT_DESC":"JNSN|PHONE.PHONE_LAST_5=10796"},{"FEAT_ID":16,"FEAT_DESC":"JNSN|SSN=3251"}],"PHONE":[{"FEAT_ID":5,"FEAT_DESC":"225-671-0796"}],"PHONE_KEY":[{"FEAT_ID":21,"FEAT_DESC":"2256710796"}],"SEARCH_KEY":[{"FEAT_ID":22,"FEAT_DESC":"LOGIN_ID:FLAVORH|"},{"FEAT_ID":23,"FEAT_DESC":"SSN:3251|80804|"}],"SSN":[{"FEAT_ID":6,"FEAT_DESC":"053-39-3251"}]},"FEATURE_SCORES":{"ACCT_NUM":[{"INBOUND_FEAT_ID":8,"INBOUND_FEAT":"5534202208773608","INBOUND_FEAT_USAGE_TYPE":"CC","CANDIDATE_FEAT_ID":8,"CANDIDATE_FEAT":"5534202208773608","CANDIDATE_FEAT_USAGE_TYPE":"CC","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"F1"}],"ADDRESS":[{"INBOUND_FEAT_ID":4,"INBOUND_FEAT":"772 Armstrong RD Delhi LA 71232","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":4,"CANDIDATE_FEAT":"772 Armstrong RD Delhi LA 71232","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"FF"}],"DOB":[{"INBOUND_FEAT_ID":2,"INBOUND_FEAT":"4/8/1983","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":100001,"CANDIDATE_FEAT":"4/8/1985","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":88,"SCORE_BUCKET":"PLAUSIBLE","SCORE_BEHAVIOR":"FMES"}],"GENDER":[{"INBOUND_FEAT_ID":3,"INBOUND_FEAT":"F","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":3,"CANDIDATE_FEAT":"F","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"FVME"}],"LOGIN_ID":[{"INBOUND_FEAT_ID":7,"INBOUND_FEAT":"flavorh","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":7,"CANDIDATE_FEAT":"flavorh","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"F1"}],"NAME":[{"INBOUND_FEAT_ID":1,"INBOUND_FEAT":"JOHNSON","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":1,"CANDIDATE_FEAT":"JOHNSON","CANDIDATE_FEAT_USAGE_TYPE":"","GNR_FN":100,"GNR_SN":100,"GNR_GN":70,"GENERATION_MATCH":-1,"GNR_ON":-1,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"NAME"}],"PHONE":[{"INBOUND_FEAT_ID":5,"INBOUND_FEAT":"225-671-0796","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":5,"CANDIDATE_FEAT":"225-671-0796","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"FF"}],"SSN":[{"INBOUND_FEAT_ID":6,"INBOUND_FEAT":"053-39-3251","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":6,"CANDIDATE_FEAT":"053-39-3251","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"F1ES"}]}}},{"INTERNAL_ID":100001,"ENTITY_ID":1,"FOCUS_RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":"111"}],"MATCH_INFO":{"WHY_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","WHY_ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","MATCH_LEVEL_CODE":"RESOLVED","CANDIDATE_KEYS":{"ACCT_NUM":[{"FEAT_ID":8,"FEAT_DESC":"5534202208773608"}],"ADDR_KEY":[{"FEAT_ID":17,"FEAT_DESC":"772|ARMSTRNK||TL"},{"FEAT_ID":18,"FEAT_DESC":"772|ARMSTRNK||71232"}],"ID_KEY":[{"FEAT_ID":19,"FEAT_DESC":"ACCT_NUM=5534202208773608"},{"FEAT_ID":20,"FEAT_DESC":"SSN=053-39-3251"}],"LOGIN_ID":[{"FEAT_ID":7,"FEAT_DESC":"flavorh"}],"NAME_KEY":[{"FEAT_ID":9,"FEAT_DESC":"JNSN|DOB.MMDD_HASH=0804"},{"FEAT_ID":11,"FEAT_DESC":"JNSN"},{"FEAT_ID":12,"FEAT_DESC":"JNSN|ADDRESS.CITY_STD=TL"},{"FEAT_ID":13,"FEAT_DESC":"JNSN|DOB=80804"},{"FEAT_ID":14,"FEAT_DESC":"JNSN|POST=71232"},{"FEAT_ID":15,"FEAT_DESC":"JNSN|PHONE.PHONE_LAST_5=10796"},{"FEAT_ID":16,"FEAT_DESC":"JNSN|SSN=3251"}],"PHONE":[{"FEAT_ID":5,"FEAT_DESC":"225-671-0796"}],"PHONE_KEY":[{"FEAT_ID":21,"FEAT_DESC":"2256710796"}],"SEARCH_KEY":[{"FEAT_ID":22,"FEAT_DESC":"LOGIN_ID:FLAVORH|"},{"FEAT_ID":23,"FEAT_DESC":"SSN:3251|80804|"}],"SSN":[{"FEAT_ID":6,"FEAT_DESC":"053-39-3251"}]},"FEATURE_SCORES":{"ACCT_NUM":[{"INBOUND_FEAT_ID":8,"INBOUND_FEAT":"5534202208773608","INBOUND_FEAT_USAGE_TYPE":"CC","CANDIDATE_FEAT_ID":8,"CANDIDATE_FEAT":"5534202208773608","CANDIDATE_FEAT_USAGE_TYPE":"CC","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"F1"}],"ADDRESS":[{"INBOUND_FEAT_ID":4,"INBOUND_FEAT":"772 Armstrong RD Delhi LA 71232","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":4,"CANDIDATE_FEAT":"772 Armstrong RD Delhi LA 71232","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"FF"}],"DOB":[{"INBOUND_FEAT_ID":100001,"INBOUND_FEAT":"4/8/1985","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":2,"CANDIDATE_FEAT":"4/8/1983","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":88,"SCORE_BUCKET":"PLAUSIBLE","SCORE_BEHAVIOR":"FMES"}],"GENDER":[{"INBOUND_FEAT_ID":3,"INBOUND_FEAT":"F","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":3,"CANDIDATE_FEAT":"F","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"FVME"}],"LOGIN_ID":[{"INBOUND_FEAT_ID":7,"INBOUND_FEAT":"flavorh","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":7,"CANDIDATE_FEAT":"flavorh","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"F1"}],"NAME":[{"INBOUND_FEAT_ID":1,"INBOUND_FEAT":"JOHNSON","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":1,"CANDIDATE_FEAT":"JOHNSON","CANDIDATE_FEAT_USAGE_TYPE":"","GNR_FN":100,"GNR_SN":100,"GNR_GN":70,"GENERATION_MATCH":-1,"GNR_ON":-1,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"NAME"}],"PHONE":[{"INBOUND_FEAT_ID":5,"INBOUND_FEAT":"225-671-0796","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":5,"CANDIDATE_FEAT":"225-671-0796","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"FF"}],"SSN":[{"INBOUND_FEAT_ID":6,"INBOUND_FEAT":"053-39-3251","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":6,"CANDIDATE_FEAT":"053-39-3251","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"F1ES"}]}}}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1,"ENTITY_NAME":"JOHNSON","FEATURES":{"ACCT_NUM":[{"FEAT_DESC":"5534202208773608","LIB_FEAT_ID":8,"USAGE_TYPE":"CC","FEAT_DESC_VALUES":[{"FEAT_DESC":"5534202208773608","LIB_FEAT_ID":8,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"ADDRESS":[{"FEAT_DESC":"772 Armstrong RD Delhi LA 71232","LIB_FEAT_ID":4,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772 Armstrong RD Delhi LA 71232","LIB_FEAT_ID":4,"USED_FOR_CAND":"N","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"ADDR_KEY":[{"FEAT_DESC":"772|ARMSTRNK||71232","LIB_FEAT_ID":18,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772|ARMSTRNK||71232","LIB_FEAT_ID":18,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"772|ARMSTRNK||TL","LIB_FEAT_ID":17,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772|ARMSTRNK||TL","LIB_FEAT_ID":17,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"DOB":[{"FEAT_DESC":"4/8/1983","LIB_FEAT_ID":2,"FEAT_DESC_VALUES":[{"FEAT_DESC":"4/8/1983","LIB_FEAT_ID":2,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"4/8/1985","LIB_FEAT_ID":100001,"FEAT_DESC_VALUES":[{"FEAT_DESC":"4/8/1985","LIB_FEAT_ID":100001,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"GENDER":[{"FEAT_DESC":"F","LIB_FEAT_ID":3,"FEAT_DESC_VALUES":[{"FEAT_DESC":"F","LIB_FEAT_ID":3,"USED_FOR_CAND":"N","USED_FOR_SCORING":"Y","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"ID_KEY":[{"FEAT_DESC":"ACCT_NUM=5534202208773608","LIB_FEAT_ID":19,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ACCT_NUM=5534202208773608","LIB_FEAT_ID":19,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"SSN=053-39-3251","LIB_FEAT_ID":20,"FEAT_DESC_VALUES":[{"FEAT_DESC":"SSN=053-39-3251","LIB_FEAT_ID":20,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"LOGIN_ID":[{"FEAT_DESC":"flavorh","LIB_FEAT_ID":7,"FEAT_DESC_VALUES":[{"FEAT_DESC":"flavorh","LIB_FEAT_ID":7,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"NAME":[{"FEAT_DESC":"JOHNSON","LIB_FEAT_ID":1,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JOHNSON","LIB_FEAT_ID":1,"USED_FOR_CAND":"N","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"NAME_KEY":[{"FEAT_DESC":"JNSN","LIB_FEAT_ID":11,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN","LIB_FEAT_ID":11,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|ADDRESS.CITY_STD=TL","LIB_FEAT_ID":12,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|ADDRESS.CITY_STD=TL","LIB_FEAT_ID":12,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|DOB.MMDD_HASH=0804","LIB_FEAT_ID":9,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|DOB.MMDD_HASH=0804","LIB_FEAT_ID":9,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|DOB.MMYY_HASH=0483","LIB_FEAT_ID":10,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|DOB.MMYY_HASH=0483","LIB_FEAT_ID":10,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|DOB.MMYY_HASH=0485","LIB_FEAT_ID":100002,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|DOB.MMYY_HASH=0485","LIB_FEAT_ID":100002,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|DOB=80804","LIB_FEAT_ID":13,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|DOB=80804","LIB_FEAT_ID":13,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|PHONE.PHONE_LAST_5=10796","LIB_FEAT_ID":15,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|PHONE.PHONE_LAST_5=10796","LIB_FEAT_ID":15,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|POST=71232","LIB_FEAT_ID":14,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|POST=71232","LIB_FEAT_ID":14,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|SSN=3251","LIB_FEAT_ID":16,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|SSN=3251","LIB_FEAT_ID":16,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"PHONE":[{"FEAT_DESC":"225-671-0796","LIB_FEAT_ID":5,"FEAT_DESC_VALUES":[{"FEAT_DESC":"225-671-0796","LIB_FEAT_ID":5,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"PHONE_KEY":[{"FEAT_DESC":"2256710796","LIB_FEAT_ID":21,"FEAT_DESC_VALUES":[{"FEAT_DESC":"2256710796","LIB_FEAT_ID":21,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"SEARCH_KEY":[{"FEAT_DESC":"LOGIN_ID:FLAVORH|","LIB_FEAT_ID":22,"FEAT_DESC_VALUES":[{"FEAT_DESC":"LOGIN_ID:FLAVORH|","LIB_FEAT_ID":22,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"SSN:3251|80804|","LIB_FEAT_ID":23,"FEAT_DESC_VALUES":[{"FEAT_DESC":"SSN:3251|80804|","LIB_FEAT_ID":23,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"SSN":[{"FEAT_DESC":"053-39-3251","LIB_FEAT_ID":6,"FEAT_DESC_VALUES":[{"FEAT_DESC":"053-39-3251","LIB_FEAT_ID":6,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}]},"RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":6,"FIRST_SEEN_DT":"2022-12-06 16:02:35.306","LAST_SEEN_DT":"2022-12-06 16:02:36.083"}],"LAST_SEEN_DT":"2022-12-06 16:02:36.083","RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":"111","ENTITY_TYPE":"TEST","INTERNAL_ID":100001,"ENTITY_KEY":"A6C927986DF7329D1D2CDE0E8F34328AE640FB7E","ENTITY_DESC":"JOHNSON","MATCH_KEY":"","MATCH_LEVEL":0,"MATCH_LEVEL_CODE":"","ERRULE_CODE":"","LAST_SEEN_DT":"2022-12-06 16:02:36.083","FEATURES":[{"LIB_FEAT_ID":1},{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":4},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":6},{"LIB_FEAT_ID":7},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":9},{"LIB_FEAT_ID":11},{"LIB_FEAT_ID":12},{"LIB_FEAT_ID":13},{"LIB_FEAT_ID":14},{"LIB_FEAT_ID":15},{"LIB_FEAT_ID":16},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":18},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":20},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":22},{"LIB_FEAT_ID":23},{"LIB_FEAT_ID":100001},{"LIB_FEAT_ID":100002}]},{"DATA_SOURCE":"TEST","RECORD_ID":"444","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"C6063D4396612FBA7324DB0739273BA1FE815C43","ENTITY_DESC":"JOHNSON","MATCH_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","MATCH_LEVEL":1,"MATCH_LEVEL_CODE":"RESOLVED","ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","LAST_SEEN_DT":"2022-12-06 16:02:35.572","FEATURES":[{"LIB_FEAT_ID":1},{"LIB_FEAT_ID":2},{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":4},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":6},{"LIB_FEAT_ID":7},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":9},{"LIB_FEAT_ID":10},{"LIB_FEAT_ID":11},{"LIB_FEAT_ID":12},{"LIB_FEAT_ID":13},{"LIB_FEAT_ID":14},{"LIB_FEAT_ID":15},{"LIB_FEAT_ID":16},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":18},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":20},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":22},{"LIB_FEAT_ID":23}]},{"DATA_SOURCE":"TEST","RECORD_ID":"555","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"C6063D4396612FBA7324DB0739273BA1FE815C43","ENTITY_DESC":"JOHNSON","MATCH_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","MATCH_LEVEL":1,"MATCH_LEVEL_CODE":"RESOLVED","ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","LAST_SEEN_DT":"2022-12-06 16:02:35.575","FEATURES":[{"LIB_FEAT_ID":1},{"LIB_FEAT_ID":2},{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":4},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":6},{"LIB_FEAT_ID":7},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":9},{"LIB_FEAT_ID":10},{"LIB_FEAT_ID":11},{"LIB_FEAT_ID":12},{"LIB_FEAT_ID":13},{"LIB_FEAT_ID":14},{"LIB_FEAT_ID":15},{"LIB_FEAT_ID":16},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":18},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":20},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":22},{"LIB_FEAT_ID":23}]},{"DATA_SOURCE":"TEST","RECORD_ID":"666","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"C6063D4396612FBA7324DB0739273BA1FE815C43","ENTITY_DESC":"JOHNSON","MATCH_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","MATCH_LEVEL":1,"MATCH_LEVEL_CODE":"RESOLVED","ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","LAST_SEEN_DT":"2022-12-06 16:02:35.579","FEATURES":[{"LIB_FEAT_ID":1},{"LIB_FEAT_ID":2},{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":4},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":6},{"LIB_FEAT_ID":7},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":9},{"LIB_FEAT_ID":10},{"LIB_FEAT_ID":11},{"LIB_FEAT_ID":12},{"LIB_FEAT_ID":13},{"LIB_FEAT_ID":14},{"LIB_FEAT_ID":15},{"LIB_FEAT_ID":16},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":18},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":20},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":22},{"LIB_FEAT_ID":23}]},{"DATA_SOURCE":"TEST","RECORD_ID":"777","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"C6063D4396612FBA7324DB0739273BA1FE815C43","ENTITY_DESC":"JOHNSON","MATCH_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","MATCH_LEVEL":1,"MATCH_LEVEL_CODE":"RESOLVED","ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","LAST_SEEN_DT":"2022-12-06 16:02:35.582","FEATURES":[{"LIB_FEAT_ID":1},{"LIB_FEAT_ID":2},{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":4},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":6},{"LIB_FEAT_ID":7},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":9},{"LIB_FEAT_ID":10},{"LIB_FEAT_ID":11},{"LIB_FEAT_ID":12},{"LIB_FEAT_ID":13},{"LIB_FEAT_ID":14},{"LIB_FEAT_ID":15},{"LIB_FEAT_ID":16},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":18},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":20},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":22},{"LIB_FEAT_ID":23}]},{"DATA_SOURCE":"TEST","RECORD_ID":"FCCE9793DAAD23159DBCCEB97FF2745B92CE7919","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"C6063D4396612FBA7324DB0739273BA1FE815C43","ENTITY_DESC":"JOHNSON","MATCH_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","MATCH_LEVEL":1,"MATCH_LEVEL_CODE":"RESOLVED","ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","LAST_SEEN_DT":"2022-12-06 16:02:35.432","FEATURES":[{"LIB_FEAT_ID":1},{"LIB_FEAT_ID":2},{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":4},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":6},{"LIB_FEAT_ID":7},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":9},{"LIB_FEAT_ID":10},{"LIB_FEAT_ID":11},{"LIB_FEAT_ID":12},{"LIB_FEAT_ID":13},{"LIB_FEAT_ID":14},{"LIB_FEAT_ID":15},{"LIB_FEAT_ID":16},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":18},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":20},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":22},{"LIB_FEAT_ID":23}]}]},"RELATED_ENTITIES":[{"ENTITY_ID":2,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0,"ENTITY_NAME":"OCEANGUY","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-12-06 16:02:35.373","LAST_SEEN_DT":"2022-12-06 16:02:35.373"}],"LAST_SEEN_DT":"2022-12-06 16:02:35.373"},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0,"ENTITY_NAME":"Smith","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-12-06 16:02:35.436","LAST_SEEN_DT":"2022-12-06 16:02:35.436"}],"LAST_SEEN_DT":"2022-12-06 16:02:35.436"}]}]}`
*/
func (client *G2engine) WhyEntityByEntityID(ctx context.Context, entityID int64) (string, error) {
	if client.isTrace {
		client.traceEntry(145, entityID)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"entityID": strconv.FormatInt(entityID, 10),
			}
			client.notify(ctx, 8069, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(146, entityID, client.WhyEntityByEntityIDResult, err, time.Since(entryTime))
	}
	return client.WhyEntityByEntityIDResult, err
}

/*
The WhyEntityByEntityID_V2 method explains why records belong to their resolved entities.
It extends WhyEntityByEntityID() by adding output control flags.

Input
  - ctx: A context to control lifecycle.
  - entityID: The unique identifier of an entity for the starting entity of the search path.
  - flags: Flags used to control information returned.

Output
  - A JSON document.
    Example: `{"WHY_RESULTS":[{"INTERNAL_ID":1,"ENTITY_ID":1,"FOCUS_RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":"444"},{"DATA_SOURCE":"TEST","RECORD_ID":"555"},{"DATA_SOURCE":"TEST","RECORD_ID":"666"},{"DATA_SOURCE":"TEST","RECORD_ID":"777"},{"DATA_SOURCE":"TEST","RECORD_ID":"FCCE9793DAAD23159DBCCEB97FF2745B92CE7919"}],"MATCH_INFO":{"WHY_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","WHY_ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","MATCH_LEVEL_CODE":"RESOLVED"}},{"INTERNAL_ID":100001,"ENTITY_ID":1,"FOCUS_RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":"111"}],"MATCH_INFO":{"WHY_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","WHY_ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","MATCH_LEVEL_CODE":"RESOLVED"}}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1}}]}`
*/
func (client *G2engine) WhyEntityByEntityID_V2(ctx context.Context, entityID int64, flags int64) (string, error) {
	if client.isTrace {
		client.traceEntry(147, entityID, flags)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"entityID": strconv.FormatInt(entityID, 10),
			}
			client.notify(ctx, 8070, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(148, entityID, flags, client.WhyEntityByEntityID_V2Result, err, time.Since(entryTime))
	}
	return client.WhyEntityByEntityID_V2Result, err
}

/*
The WhyEntityByRecordID method explains why records belong to their resolved entities.
To control output, use WhyEntityByRecordID_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - recordID: The unique identifier within the records of the same data source.

Output
  - A JSON document.
    Example: `{"WHY_RESULTS":[{"INTERNAL_ID":1,"ENTITY_ID":1,"FOCUS_RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":"444"},{"DATA_SOURCE":"TEST","RECORD_ID":"555"},{"DATA_SOURCE":"TEST","RECORD_ID":"666"},{"DATA_SOURCE":"TEST","RECORD_ID":"777"},{"DATA_SOURCE":"TEST","RECORD_ID":"FCCE9793DAAD23159DBCCEB97FF2745B92CE7919"}],"MATCH_INFO":{"WHY_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","WHY_ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","MATCH_LEVEL_CODE":"RESOLVED","CANDIDATE_KEYS":{"ACCT_NUM":[{"FEAT_ID":8,"FEAT_DESC":"5534202208773608"}],"ADDR_KEY":[{"FEAT_ID":17,"FEAT_DESC":"772|ARMSTRNK||TL"},{"FEAT_ID":18,"FEAT_DESC":"772|ARMSTRNK||71232"}],"ID_KEY":[{"FEAT_ID":19,"FEAT_DESC":"ACCT_NUM=5534202208773608"},{"FEAT_ID":20,"FEAT_DESC":"SSN=053-39-3251"}],"LOGIN_ID":[{"FEAT_ID":7,"FEAT_DESC":"flavorh"}],"NAME_KEY":[{"FEAT_ID":9,"FEAT_DESC":"JNSN|DOB.MMDD_HASH=0804"},{"FEAT_ID":11,"FEAT_DESC":"JNSN"},{"FEAT_ID":12,"FEAT_DESC":"JNSN|ADDRESS.CITY_STD=TL"},{"FEAT_ID":13,"FEAT_DESC":"JNSN|DOB=80804"},{"FEAT_ID":14,"FEAT_DESC":"JNSN|POST=71232"},{"FEAT_ID":15,"FEAT_DESC":"JNSN|PHONE.PHONE_LAST_5=10796"},{"FEAT_ID":16,"FEAT_DESC":"JNSN|SSN=3251"}],"PHONE":[{"FEAT_ID":5,"FEAT_DESC":"225-671-0796"}],"PHONE_KEY":[{"FEAT_ID":21,"FEAT_DESC":"2256710796"}],"SEARCH_KEY":[{"FEAT_ID":22,"FEAT_DESC":"LOGIN_ID:FLAVORH|"},{"FEAT_ID":23,"FEAT_DESC":"SSN:3251|80804|"}],"SSN":[{"FEAT_ID":6,"FEAT_DESC":"053-39-3251"}]},"FEATURE_SCORES":{"ACCT_NUM":[{"INBOUND_FEAT_ID":8,"INBOUND_FEAT":"5534202208773608","INBOUND_FEAT_USAGE_TYPE":"CC","CANDIDATE_FEAT_ID":8,"CANDIDATE_FEAT":"5534202208773608","CANDIDATE_FEAT_USAGE_TYPE":"CC","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"F1"}],"ADDRESS":[{"INBOUND_FEAT_ID":4,"INBOUND_FEAT":"772 Armstrong RD Delhi LA 71232","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":4,"CANDIDATE_FEAT":"772 Armstrong RD Delhi LA 71232","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"FF"}],"DOB":[{"INBOUND_FEAT_ID":2,"INBOUND_FEAT":"4/8/1983","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":100001,"CANDIDATE_FEAT":"4/8/1985","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":88,"SCORE_BUCKET":"PLAUSIBLE","SCORE_BEHAVIOR":"FMES"}],"GENDER":[{"INBOUND_FEAT_ID":3,"INBOUND_FEAT":"F","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":3,"CANDIDATE_FEAT":"F","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"FVME"}],"LOGIN_ID":[{"INBOUND_FEAT_ID":7,"INBOUND_FEAT":"flavorh","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":7,"CANDIDATE_FEAT":"flavorh","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"F1"}],"NAME":[{"INBOUND_FEAT_ID":1,"INBOUND_FEAT":"JOHNSON","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":1,"CANDIDATE_FEAT":"JOHNSON","CANDIDATE_FEAT_USAGE_TYPE":"","GNR_FN":100,"GNR_SN":100,"GNR_GN":70,"GENERATION_MATCH":-1,"GNR_ON":-1,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"NAME"}],"PHONE":[{"INBOUND_FEAT_ID":5,"INBOUND_FEAT":"225-671-0796","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":5,"CANDIDATE_FEAT":"225-671-0796","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"FF"}],"SSN":[{"INBOUND_FEAT_ID":6,"INBOUND_FEAT":"053-39-3251","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":6,"CANDIDATE_FEAT":"053-39-3251","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"F1ES"}]}}},{"INTERNAL_ID":100001,"ENTITY_ID":1,"FOCUS_RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":"111"}],"MATCH_INFO":{"WHY_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","WHY_ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","MATCH_LEVEL_CODE":"RESOLVED","CANDIDATE_KEYS":{"ACCT_NUM":[{"FEAT_ID":8,"FEAT_DESC":"5534202208773608"}],"ADDR_KEY":[{"FEAT_ID":17,"FEAT_DESC":"772|ARMSTRNK||TL"},{"FEAT_ID":18,"FEAT_DESC":"772|ARMSTRNK||71232"}],"ID_KEY":[{"FEAT_ID":19,"FEAT_DESC":"ACCT_NUM=5534202208773608"},{"FEAT_ID":20,"FEAT_DESC":"SSN=053-39-3251"}],"LOGIN_ID":[{"FEAT_ID":7,"FEAT_DESC":"flavorh"}],"NAME_KEY":[{"FEAT_ID":9,"FEAT_DESC":"JNSN|DOB.MMDD_HASH=0804"},{"FEAT_ID":11,"FEAT_DESC":"JNSN"},{"FEAT_ID":12,"FEAT_DESC":"JNSN|ADDRESS.CITY_STD=TL"},{"FEAT_ID":13,"FEAT_DESC":"JNSN|DOB=80804"},{"FEAT_ID":14,"FEAT_DESC":"JNSN|POST=71232"},{"FEAT_ID":15,"FEAT_DESC":"JNSN|PHONE.PHONE_LAST_5=10796"},{"FEAT_ID":16,"FEAT_DESC":"JNSN|SSN=3251"}],"PHONE":[{"FEAT_ID":5,"FEAT_DESC":"225-671-0796"}],"PHONE_KEY":[{"FEAT_ID":21,"FEAT_DESC":"2256710796"}],"SEARCH_KEY":[{"FEAT_ID":22,"FEAT_DESC":"LOGIN_ID:FLAVORH|"},{"FEAT_ID":23,"FEAT_DESC":"SSN:3251|80804|"}],"SSN":[{"FEAT_ID":6,"FEAT_DESC":"053-39-3251"}]},"FEATURE_SCORES":{"ACCT_NUM":[{"INBOUND_FEAT_ID":8,"INBOUND_FEAT":"5534202208773608","INBOUND_FEAT_USAGE_TYPE":"CC","CANDIDATE_FEAT_ID":8,"CANDIDATE_FEAT":"5534202208773608","CANDIDATE_FEAT_USAGE_TYPE":"CC","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"F1"}],"ADDRESS":[{"INBOUND_FEAT_ID":4,"INBOUND_FEAT":"772 Armstrong RD Delhi LA 71232","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":4,"CANDIDATE_FEAT":"772 Armstrong RD Delhi LA 71232","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"FF"}],"DOB":[{"INBOUND_FEAT_ID":100001,"INBOUND_FEAT":"4/8/1985","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":2,"CANDIDATE_FEAT":"4/8/1983","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":88,"SCORE_BUCKET":"PLAUSIBLE","SCORE_BEHAVIOR":"FMES"}],"GENDER":[{"INBOUND_FEAT_ID":3,"INBOUND_FEAT":"F","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":3,"CANDIDATE_FEAT":"F","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"FVME"}],"LOGIN_ID":[{"INBOUND_FEAT_ID":7,"INBOUND_FEAT":"flavorh","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":7,"CANDIDATE_FEAT":"flavorh","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"F1"}],"NAME":[{"INBOUND_FEAT_ID":1,"INBOUND_FEAT":"JOHNSON","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":1,"CANDIDATE_FEAT":"JOHNSON","CANDIDATE_FEAT_USAGE_TYPE":"","GNR_FN":100,"GNR_SN":100,"GNR_GN":70,"GENERATION_MATCH":-1,"GNR_ON":-1,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"NAME"}],"PHONE":[{"INBOUND_FEAT_ID":5,"INBOUND_FEAT":"225-671-0796","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":5,"CANDIDATE_FEAT":"225-671-0796","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"FF"}],"SSN":[{"INBOUND_FEAT_ID":6,"INBOUND_FEAT":"053-39-3251","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":6,"CANDIDATE_FEAT":"053-39-3251","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"F1ES"}]}}}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1,"ENTITY_NAME":"JOHNSON","FEATURES":{"ACCT_NUM":[{"FEAT_DESC":"5534202208773608","LIB_FEAT_ID":8,"USAGE_TYPE":"CC","FEAT_DESC_VALUES":[{"FEAT_DESC":"5534202208773608","LIB_FEAT_ID":8,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"ADDRESS":[{"FEAT_DESC":"772 Armstrong RD Delhi LA 71232","LIB_FEAT_ID":4,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772 Armstrong RD Delhi LA 71232","LIB_FEAT_ID":4,"USED_FOR_CAND":"N","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"ADDR_KEY":[{"FEAT_DESC":"772|ARMSTRNK||71232","LIB_FEAT_ID":18,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772|ARMSTRNK||71232","LIB_FEAT_ID":18,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"772|ARMSTRNK||TL","LIB_FEAT_ID":17,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772|ARMSTRNK||TL","LIB_FEAT_ID":17,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"DOB":[{"FEAT_DESC":"4/8/1983","LIB_FEAT_ID":2,"FEAT_DESC_VALUES":[{"FEAT_DESC":"4/8/1983","LIB_FEAT_ID":2,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"4/8/1985","LIB_FEAT_ID":100001,"FEAT_DESC_VALUES":[{"FEAT_DESC":"4/8/1985","LIB_FEAT_ID":100001,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"GENDER":[{"FEAT_DESC":"F","LIB_FEAT_ID":3,"FEAT_DESC_VALUES":[{"FEAT_DESC":"F","LIB_FEAT_ID":3,"USED_FOR_CAND":"N","USED_FOR_SCORING":"Y","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"ID_KEY":[{"FEAT_DESC":"ACCT_NUM=5534202208773608","LIB_FEAT_ID":19,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ACCT_NUM=5534202208773608","LIB_FEAT_ID":19,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"SSN=053-39-3251","LIB_FEAT_ID":20,"FEAT_DESC_VALUES":[{"FEAT_DESC":"SSN=053-39-3251","LIB_FEAT_ID":20,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"LOGIN_ID":[{"FEAT_DESC":"flavorh","LIB_FEAT_ID":7,"FEAT_DESC_VALUES":[{"FEAT_DESC":"flavorh","LIB_FEAT_ID":7,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"NAME":[{"FEAT_DESC":"JOHNSON","LIB_FEAT_ID":1,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JOHNSON","LIB_FEAT_ID":1,"USED_FOR_CAND":"N","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"NAME_KEY":[{"FEAT_DESC":"JNSN","LIB_FEAT_ID":11,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN","LIB_FEAT_ID":11,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|ADDRESS.CITY_STD=TL","LIB_FEAT_ID":12,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|ADDRESS.CITY_STD=TL","LIB_FEAT_ID":12,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|DOB.MMDD_HASH=0804","LIB_FEAT_ID":9,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|DOB.MMDD_HASH=0804","LIB_FEAT_ID":9,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|DOB.MMYY_HASH=0483","LIB_FEAT_ID":10,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|DOB.MMYY_HASH=0483","LIB_FEAT_ID":10,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|DOB.MMYY_HASH=0485","LIB_FEAT_ID":100002,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|DOB.MMYY_HASH=0485","LIB_FEAT_ID":100002,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|DOB=80804","LIB_FEAT_ID":13,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|DOB=80804","LIB_FEAT_ID":13,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|PHONE.PHONE_LAST_5=10796","LIB_FEAT_ID":15,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|PHONE.PHONE_LAST_5=10796","LIB_FEAT_ID":15,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|POST=71232","LIB_FEAT_ID":14,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|POST=71232","LIB_FEAT_ID":14,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|SSN=3251","LIB_FEAT_ID":16,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|SSN=3251","LIB_FEAT_ID":16,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"PHONE":[{"FEAT_DESC":"225-671-0796","LIB_FEAT_ID":5,"FEAT_DESC_VALUES":[{"FEAT_DESC":"225-671-0796","LIB_FEAT_ID":5,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"PHONE_KEY":[{"FEAT_DESC":"2256710796","LIB_FEAT_ID":21,"FEAT_DESC_VALUES":[{"FEAT_DESC":"2256710796","LIB_FEAT_ID":21,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"SEARCH_KEY":[{"FEAT_DESC":"LOGIN_ID:FLAVORH|","LIB_FEAT_ID":22,"FEAT_DESC_VALUES":[{"FEAT_DESC":"LOGIN_ID:FLAVORH|","LIB_FEAT_ID":22,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"SSN:3251|80804|","LIB_FEAT_ID":23,"FEAT_DESC_VALUES":[{"FEAT_DESC":"SSN:3251|80804|","LIB_FEAT_ID":23,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"SSN":[{"FEAT_DESC":"053-39-3251","LIB_FEAT_ID":6,"FEAT_DESC_VALUES":[{"FEAT_DESC":"053-39-3251","LIB_FEAT_ID":6,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}]},"RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":6,"FIRST_SEEN_DT":"2022-12-06 16:09:05.626","LAST_SEEN_DT":"2022-12-06 16:09:06.399"}],"LAST_SEEN_DT":"2022-12-06 16:09:06.399","RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":"111","ENTITY_TYPE":"TEST","INTERNAL_ID":100001,"ENTITY_KEY":"A6C927986DF7329D1D2CDE0E8F34328AE640FB7E","ENTITY_DESC":"JOHNSON","MATCH_KEY":"","MATCH_LEVEL":0,"MATCH_LEVEL_CODE":"","ERRULE_CODE":"","LAST_SEEN_DT":"2022-12-06 16:09:06.399","FEATURES":[{"LIB_FEAT_ID":1},{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":4},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":6},{"LIB_FEAT_ID":7},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":9},{"LIB_FEAT_ID":11},{"LIB_FEAT_ID":12},{"LIB_FEAT_ID":13},{"LIB_FEAT_ID":14},{"LIB_FEAT_ID":15},{"LIB_FEAT_ID":16},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":18},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":20},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":22},{"LIB_FEAT_ID":23},{"LIB_FEAT_ID":100001},{"LIB_FEAT_ID":100002}]},{"DATA_SOURCE":"TEST","RECORD_ID":"444","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"C6063D4396612FBA7324DB0739273BA1FE815C43","ENTITY_DESC":"JOHNSON","MATCH_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","MATCH_LEVEL":1,"MATCH_LEVEL_CODE":"RESOLVED","ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","LAST_SEEN_DT":"2022-12-06 16:09:05.954","FEATURES":[{"LIB_FEAT_ID":1},{"LIB_FEAT_ID":2},{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":4},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":6},{"LIB_FEAT_ID":7},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":9},{"LIB_FEAT_ID":10},{"LIB_FEAT_ID":11},{"LIB_FEAT_ID":12},{"LIB_FEAT_ID":13},{"LIB_FEAT_ID":14},{"LIB_FEAT_ID":15},{"LIB_FEAT_ID":16},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":18},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":20},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":22},{"LIB_FEAT_ID":23}]},{"DATA_SOURCE":"TEST","RECORD_ID":"555","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"C6063D4396612FBA7324DB0739273BA1FE815C43","ENTITY_DESC":"JOHNSON","MATCH_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","MATCH_LEVEL":1,"MATCH_LEVEL_CODE":"RESOLVED","ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","LAST_SEEN_DT":"2022-12-06 16:09:05.957","FEATURES":[{"LIB_FEAT_ID":1},{"LIB_FEAT_ID":2},{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":4},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":6},{"LIB_FEAT_ID":7},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":9},{"LIB_FEAT_ID":10},{"LIB_FEAT_ID":11},{"LIB_FEAT_ID":12},{"LIB_FEAT_ID":13},{"LIB_FEAT_ID":14},{"LIB_FEAT_ID":15},{"LIB_FEAT_ID":16},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":18},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":20},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":22},{"LIB_FEAT_ID":23}]},{"DATA_SOURCE":"TEST","RECORD_ID":"666","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"C6063D4396612FBA7324DB0739273BA1FE815C43","ENTITY_DESC":"JOHNSON","MATCH_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","MATCH_LEVEL":1,"MATCH_LEVEL_CODE":"RESOLVED","ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","LAST_SEEN_DT":"2022-12-06 16:09:05.960","FEATURES":[{"LIB_FEAT_ID":1},{"LIB_FEAT_ID":2},{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":4},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":6},{"LIB_FEAT_ID":7},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":9},{"LIB_FEAT_ID":10},{"LIB_FEAT_ID":11},{"LIB_FEAT_ID":12},{"LIB_FEAT_ID":13},{"LIB_FEAT_ID":14},{"LIB_FEAT_ID":15},{"LIB_FEAT_ID":16},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":18},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":20},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":22},{"LIB_FEAT_ID":23}]},{"DATA_SOURCE":"TEST","RECORD_ID":"777","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"C6063D4396612FBA7324DB0739273BA1FE815C43","ENTITY_DESC":"JOHNSON","MATCH_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","MATCH_LEVEL":1,"MATCH_LEVEL_CODE":"RESOLVED","ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","LAST_SEEN_DT":"2022-12-06 16:09:05.963","FEATURES":[{"LIB_FEAT_ID":1},{"LIB_FEAT_ID":2},{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":4},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":6},{"LIB_FEAT_ID":7},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":9},{"LIB_FEAT_ID":10},{"LIB_FEAT_ID":11},{"LIB_FEAT_ID":12},{"LIB_FEAT_ID":13},{"LIB_FEAT_ID":14},{"LIB_FEAT_ID":15},{"LIB_FEAT_ID":16},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":18},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":20},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":22},{"LIB_FEAT_ID":23}]},{"DATA_SOURCE":"TEST","RECORD_ID":"FCCE9793DAAD23159DBCCEB97FF2745B92CE7919","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"C6063D4396612FBA7324DB0739273BA1FE815C43","ENTITY_DESC":"JOHNSON","MATCH_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","MATCH_LEVEL":1,"MATCH_LEVEL_CODE":"RESOLVED","ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","LAST_SEEN_DT":"2022-12-06 16:09:05.789","FEATURES":[{"LIB_FEAT_ID":1},{"LIB_FEAT_ID":2},{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":4},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":6},{"LIB_FEAT_ID":7},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":9},{"LIB_FEAT_ID":10},{"LIB_FEAT_ID":11},{"LIB_FEAT_ID":12},{"LIB_FEAT_ID":13},{"LIB_FEAT_ID":14},{"LIB_FEAT_ID":15},{"LIB_FEAT_ID":16},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":18},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":20},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":22},{"LIB_FEAT_ID":23}]}]},"RELATED_ENTITIES":[{"ENTITY_ID":2,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0,"ENTITY_NAME":"OCEANGUY","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-12-06 16:09:05.724","LAST_SEEN_DT":"2022-12-06 16:09:05.724"}],"LAST_SEEN_DT":"2022-12-06 16:09:05.724"},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0,"ENTITY_NAME":"Smith","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-12-06 16:09:05.794","LAST_SEEN_DT":"2022-12-06 16:09:05.794"}],"LAST_SEEN_DT":"2022-12-06 16:09:05.794"}]}]}`
*/
func (client *G2engine) WhyEntityByRecordID(ctx context.Context, dataSourceCode string, recordID string) (string, error) {
	if client.isTrace {
		client.traceEntry(149, dataSourceCode, recordID)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"dataSourceCode": dataSourceCode,
				"recordID":       recordID,
			}
			client.notify(ctx, 8071, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(150, dataSourceCode, recordID, client.WhyEntityByRecordIDResult, err, time.Since(entryTime))
	}
	return client.WhyEntityByRecordIDResult, err
}

/*
The WhyEntityByRecordID_V2 method explains why records belong to their resolved entities.
It extends WhyEntityByRecordID() by adding output control flags.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - recordID: The unique identifier within the records of the same data source.
  - flags: Flags used to control information returned.

Output
  - A JSON document.
    Example: `{"WHY_RESULTS":[{"INTERNAL_ID":1,"ENTITY_ID":1,"FOCUS_RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":"444"},{"DATA_SOURCE":"TEST","RECORD_ID":"555"},{"DATA_SOURCE":"TEST","RECORD_ID":"666"},{"DATA_SOURCE":"TEST","RECORD_ID":"777"},{"DATA_SOURCE":"TEST","RECORD_ID":"FCCE9793DAAD23159DBCCEB97FF2745B92CE7919"}],"MATCH_INFO":{"WHY_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","WHY_ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","MATCH_LEVEL_CODE":"RESOLVED"}},{"INTERNAL_ID":100001,"ENTITY_ID":1,"FOCUS_RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":"111"}],"MATCH_INFO":{"WHY_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","WHY_ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","MATCH_LEVEL_CODE":"RESOLVED"}}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1}}]}`
*/
func (client *G2engine) WhyEntityByRecordID_V2(ctx context.Context, dataSourceCode string, recordID string, flags int64) (string, error) {
	if client.isTrace {
		client.traceEntry(151, dataSourceCode, recordID, flags)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"dataSourceCode": dataSourceCode,
				"recordID":       recordID,
			}
			client.notify(ctx, 8072, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(152, dataSourceCode, recordID, flags, client.WhyEntityByRecordID_V2Result, err, time.Since(entryTime))
	}
	return client.WhyEntityByRecordID_V2Result, err
}

/*
The WhyRecords method explains why records belong to their resolved entities.
To control output, use WhyRecords_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode1: Identifies the provenance of the data.
  - recordID1: The unique identifier within the records of the same data source.
  - dataSourceCode2: Identifies the provenance of the data.
  - recordID2: The unique identifier within the records of the same data source.

Output

  - A JSON document.
    Example: `{"WHY_RESULTS":[{"INTERNAL_ID":100001,"ENTITY_ID":1,"FOCUS_RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":"111"}],"INTERNAL_ID_2":2,"ENTITY_ID_2":2,"FOCUS_RECORDS_2":[{"DATA_SOURCE":"TEST","RECORD_ID":"222"}],"MATCH_INFO":{"WHY_KEY":"+PHONE+ACCT_NUM-DOB-SSN","WHY_ERRULE_CODE":"SF1","MATCH_LEVEL_CODE":"POSSIBLY_RELATED","CANDIDATE_KEYS":{"ACCT_NUM":[{"FEAT_ID":8,"FEAT_DESC":"5534202208773608"}],"ADDR_KEY":[{"FEAT_ID":17,"FEAT_DESC":"772|ARMSTRNK||TL"}],"ID_KEY":[{"FEAT_ID":19,"FEAT_DESC":"ACCT_NUM=5534202208773608"}],"PHONE":[{"FEAT_ID":5,"FEAT_DESC":"225-671-0796"}],"PHONE_KEY":[{"FEAT_ID":21,"FEAT_DESC":"2256710796"}]},"DISCLOSED_RELATIONS":{},"FEATURE_SCORES":{"ACCT_NUM":[{"INBOUND_FEAT_ID":8,"INBOUND_FEAT":"5534202208773608","INBOUND_FEAT_USAGE_TYPE":"CC","CANDIDATE_FEAT_ID":8,"CANDIDATE_FEAT":"5534202208773608","CANDIDATE_FEAT_USAGE_TYPE":"CC","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"F1"}],"ADDRESS":[{"INBOUND_FEAT_ID":4,"INBOUND_FEAT":"772 Armstrong RD Delhi LA 71232","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":26,"CANDIDATE_FEAT":"772 Armstrong RD Delhi WI 53543","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":81,"SCORE_BUCKET":"LIKELY","SCORE_BEHAVIOR":"FF"}],"DOB":[{"INBOUND_FEAT_ID":100001,"INBOUND_FEAT":"4/8/1985","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":25,"CANDIDATE_FEAT":"6/9/1983","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":79,"SCORE_BUCKET":"NO_CHANCE","SCORE_BEHAVIOR":"FMES"}],"GENDER":[{"INBOUND_FEAT_ID":3,"INBOUND_FEAT":"F","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":3,"CANDIDATE_FEAT":"F","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"FVME"}],"LOGIN_ID":[{"INBOUND_FEAT_ID":7,"INBOUND_FEAT":"flavorh","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":28,"CANDIDATE_FEAT":"flavorh2","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":0,"SCORE_BUCKET":"NO_CHANCE","SCORE_BEHAVIOR":"F1"}],"NAME":[{"INBOUND_FEAT_ID":1,"INBOUND_FEAT":"JOHNSON","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":24,"CANDIDATE_FEAT":"OCEANGUY","CANDIDATE_FEAT_USAGE_TYPE":"","GNR_FN":33,"GNR_SN":32,"GNR_GN":70,"GENERATION_MATCH":-1,"GNR_ON":-1,"SCORE_BUCKET":"NO_CHANCE","SCORE_BEHAVIOR":"NAME"}],"PHONE":[{"INBOUND_FEAT_ID":5,"INBOUND_FEAT":"225-671-0796","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":5,"CANDIDATE_FEAT":"225-671-0796","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"FF"}],"SSN":[{"INBOUND_FEAT_ID":6,"INBOUND_FEAT":"053-39-3251","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":27,"CANDIDATE_FEAT":"153-33-5185","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":0,"SCORE_BUCKET":"NO_CHANCE","SCORE_BEHAVIOR":"F1ES"}]}}}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1,"ENTITY_NAME":"JOHNSON","FEATURES":{"ACCT_NUM":[{"FEAT_DESC":"5534202208773608","LIB_FEAT_ID":8,"USAGE_TYPE":"CC","FEAT_DESC_VALUES":[{"FEAT_DESC":"5534202208773608","LIB_FEAT_ID":8,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"ADDRESS":[{"FEAT_DESC":"772 Armstrong RD Delhi LA 71232","LIB_FEAT_ID":4,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772 Armstrong RD Delhi LA 71232","LIB_FEAT_ID":4,"USED_FOR_CAND":"N","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"ADDR_KEY":[{"FEAT_DESC":"772|ARMSTRNK||71232","LIB_FEAT_ID":18,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772|ARMSTRNK||71232","LIB_FEAT_ID":18,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"772|ARMSTRNK||TL","LIB_FEAT_ID":17,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772|ARMSTRNK||TL","LIB_FEAT_ID":17,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"DOB":[{"FEAT_DESC":"4/8/1983","LIB_FEAT_ID":2,"FEAT_DESC_VALUES":[{"FEAT_DESC":"4/8/1983","LIB_FEAT_ID":2,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"4/8/1985","LIB_FEAT_ID":100001,"FEAT_DESC_VALUES":[{"FEAT_DESC":"4/8/1985","LIB_FEAT_ID":100001,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"GENDER":[{"FEAT_DESC":"F","LIB_FEAT_ID":3,"FEAT_DESC_VALUES":[{"FEAT_DESC":"F","LIB_FEAT_ID":3,"USED_FOR_CAND":"N","USED_FOR_SCORING":"Y","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"ID_KEY":[{"FEAT_DESC":"ACCT_NUM=5534202208773608","LIB_FEAT_ID":19,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ACCT_NUM=5534202208773608","LIB_FEAT_ID":19,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"SSN=053-39-3251","LIB_FEAT_ID":20,"FEAT_DESC_VALUES":[{"FEAT_DESC":"SSN=053-39-3251","LIB_FEAT_ID":20,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"LOGIN_ID":[{"FEAT_DESC":"flavorh","LIB_FEAT_ID":7,"FEAT_DESC_VALUES":[{"FEAT_DESC":"flavorh","LIB_FEAT_ID":7,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"NAME":[{"FEAT_DESC":"JOHNSON","LIB_FEAT_ID":1,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JOHNSON","LIB_FEAT_ID":1,"USED_FOR_CAND":"N","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"NAME_KEY":[{"FEAT_DESC":"JNSN","LIB_FEAT_ID":11,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN","LIB_FEAT_ID":11,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|ADDRESS.CITY_STD=TL","LIB_FEAT_ID":12,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|ADDRESS.CITY_STD=TL","LIB_FEAT_ID":12,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|DOB.MMDD_HASH=0804","LIB_FEAT_ID":9,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|DOB.MMDD_HASH=0804","LIB_FEAT_ID":9,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|DOB.MMYY_HASH=0483","LIB_FEAT_ID":10,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|DOB.MMYY_HASH=0483","LIB_FEAT_ID":10,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|DOB.MMYY_HASH=0485","LIB_FEAT_ID":100002,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|DOB.MMYY_HASH=0485","LIB_FEAT_ID":100002,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|DOB=80804","LIB_FEAT_ID":13,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|DOB=80804","LIB_FEAT_ID":13,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|PHONE.PHONE_LAST_5=10796","LIB_FEAT_ID":15,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|PHONE.PHONE_LAST_5=10796","LIB_FEAT_ID":15,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|POST=71232","LIB_FEAT_ID":14,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|POST=71232","LIB_FEAT_ID":14,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|SSN=3251","LIB_FEAT_ID":16,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|SSN=3251","LIB_FEAT_ID":16,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"PHONE":[{"FEAT_DESC":"225-671-0796","LIB_FEAT_ID":5,"FEAT_DESC_VALUES":[{"FEAT_DESC":"225-671-0796","LIB_FEAT_ID":5,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"PHONE_KEY":[{"FEAT_DESC":"2256710796","LIB_FEAT_ID":21,"FEAT_DESC_VALUES":[{"FEAT_DESC":"2256710796","LIB_FEAT_ID":21,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"SEARCH_KEY":[{"FEAT_DESC":"LOGIN_ID:FLAVORH|","LIB_FEAT_ID":22,"FEAT_DESC_VALUES":[{"FEAT_DESC":"LOGIN_ID:FLAVORH|","LIB_FEAT_ID":22,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"SSN:3251|80804|","LIB_FEAT_ID":23,"FEAT_DESC_VALUES":[{"FEAT_DESC":"SSN:3251|80804|","LIB_FEAT_ID":23,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"SSN":[{"FEAT_DESC":"053-39-3251","LIB_FEAT_ID":6,"FEAT_DESC_VALUES":[{"FEAT_DESC":"053-39-3251","LIB_FEAT_ID":6,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}]},"RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":6,"FIRST_SEEN_DT":"2022-12-06 16:13:27.135","LAST_SEEN_DT":"2022-12-06 16:13:27.916"}],"LAST_SEEN_DT":"2022-12-06 16:13:27.916","RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":"111","ENTITY_TYPE":"TEST","INTERNAL_ID":100001,"ENTITY_KEY":"A6C927986DF7329D1D2CDE0E8F34328AE640FB7E","ENTITY_DESC":"JOHNSON","MATCH_KEY":"","MATCH_LEVEL":0,"MATCH_LEVEL_CODE":"","ERRULE_CODE":"","LAST_SEEN_DT":"2022-12-06 16:13:27.916","FEATURES":[{"LIB_FEAT_ID":1},{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":4},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":6},{"LIB_FEAT_ID":7},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":9},{"LIB_FEAT_ID":11},{"LIB_FEAT_ID":12},{"LIB_FEAT_ID":13},{"LIB_FEAT_ID":14},{"LIB_FEAT_ID":15},{"LIB_FEAT_ID":16},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":18},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":20},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":22},{"LIB_FEAT_ID":23},{"LIB_FEAT_ID":100001},{"LIB_FEAT_ID":100002}]},{"DATA_SOURCE":"TEST","RECORD_ID":"444","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"C6063D4396612FBA7324DB0739273BA1FE815C43","ENTITY_DESC":"JOHNSON","MATCH_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","MATCH_LEVEL":1,"MATCH_LEVEL_CODE":"RESOLVED","ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","LAST_SEEN_DT":"2022-12-06 16:13:27.405","FEATURES":[{"LIB_FEAT_ID":1},{"LIB_FEAT_ID":2},{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":4},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":6},{"LIB_FEAT_ID":7},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":9},{"LIB_FEAT_ID":10},{"LIB_FEAT_ID":11},{"LIB_FEAT_ID":12},{"LIB_FEAT_ID":13},{"LIB_FEAT_ID":14},{"LIB_FEAT_ID":15},{"LIB_FEAT_ID":16},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":18},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":20},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":22},{"LIB_FEAT_ID":23}]},{"DATA_SOURCE":"TEST","RECORD_ID":"555","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"C6063D4396612FBA7324DB0739273BA1FE815C43","ENTITY_DESC":"JOHNSON","MATCH_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","MATCH_LEVEL":1,"MATCH_LEVEL_CODE":"RESOLVED","ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","LAST_SEEN_DT":"2022-12-06 16:13:27.408","FEATURES":[{"LIB_FEAT_ID":1},{"LIB_FEAT_ID":2},{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":4},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":6},{"LIB_FEAT_ID":7},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":9},{"LIB_FEAT_ID":10},{"LIB_FEAT_ID":11},{"LIB_FEAT_ID":12},{"LIB_FEAT_ID":13},{"LIB_FEAT_ID":14},{"LIB_FEAT_ID":15},{"LIB_FEAT_ID":16},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":18},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":20},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":22},{"LIB_FEAT_ID":23}]},{"DATA_SOURCE":"TEST","RECORD_ID":"666","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"C6063D4396612FBA7324DB0739273BA1FE815C43","ENTITY_DESC":"JOHNSON","MATCH_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","MATCH_LEVEL":1,"MATCH_LEVEL_CODE":"RESOLVED","ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","LAST_SEEN_DT":"2022-12-06 16:13:27.411","FEATURES":[{"LIB_FEAT_ID":1},{"LIB_FEAT_ID":2},{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":4},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":6},{"LIB_FEAT_ID":7},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":9},{"LIB_FEAT_ID":10},{"LIB_FEAT_ID":11},{"LIB_FEAT_ID":12},{"LIB_FEAT_ID":13},{"LIB_FEAT_ID":14},{"LIB_FEAT_ID":15},{"LIB_FEAT_ID":16},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":18},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":20},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":22},{"LIB_FEAT_ID":23}]},{"DATA_SOURCE":"TEST","RECORD_ID":"777","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"C6063D4396612FBA7324DB0739273BA1FE815C43","ENTITY_DESC":"JOHNSON","MATCH_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","MATCH_LEVEL":1,"MATCH_LEVEL_CODE":"RESOLVED","ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","LAST_SEEN_DT":"2022-12-06 16:13:27.418","FEATURES":[{"LIB_FEAT_ID":1},{"LIB_FEAT_ID":2},{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":4},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":6},{"LIB_FEAT_ID":7},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":9},{"LIB_FEAT_ID":10},{"LIB_FEAT_ID":11},{"LIB_FEAT_ID":12},{"LIB_FEAT_ID":13},{"LIB_FEAT_ID":14},{"LIB_FEAT_ID":15},{"LIB_FEAT_ID":16},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":18},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":20},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":22},{"LIB_FEAT_ID":23}]},{"DATA_SOURCE":"TEST","RECORD_ID":"FCCE9793DAAD23159DBCCEB97FF2745B92CE7919","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"C6063D4396612FBA7324DB0739273BA1FE815C43","ENTITY_DESC":"JOHNSON","MATCH_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","MATCH_LEVEL":1,"MATCH_LEVEL_CODE":"RESOLVED","ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","LAST_SEEN_DT":"2022-12-06 16:13:27.265","FEATURES":[{"LIB_FEAT_ID":1},{"LIB_FEAT_ID":2},{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":4},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":6},{"LIB_FEAT_ID":7},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":9},{"LIB_FEAT_ID":10},{"LIB_FEAT_ID":11},{"LIB_FEAT_ID":12},{"LIB_FEAT_ID":13},{"LIB_FEAT_ID":14},{"LIB_FEAT_ID":15},{"LIB_FEAT_ID":16},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":18},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":20},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":22},{"LIB_FEAT_ID":23}]}]},"RELATED_ENTITIES":[{"ENTITY_ID":2,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0,"ENTITY_NAME":"OCEANGUY","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-12-06 16:13:27.208","LAST_SEEN_DT":"2022-12-06 16:13:27.208"}],"LAST_SEEN_DT":"2022-12-06 16:13:27.208"},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0,"ENTITY_NAME":"Smith","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-12-06 16:13:27.272","LAST_SEEN_DT":"2022-12-06 16:13:27.272"}],"LAST_SEEN_DT":"2022-12-06 16:13:27.272"}]},{"RESOLVED_ENTITY":{"ENTITY_ID":2,"ENTITY_NAME":"OCEANGUY","FEATURES":{"ACCT_NUM":[{"FEAT_DESC":"5534202208773608","LIB_FEAT_ID":8,"USAGE_TYPE":"CC","FEAT_DESC_VALUES":[{"FEAT_DESC":"5534202208773608","LIB_FEAT_ID":8,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"ADDRESS":[{"FEAT_DESC":"772 Armstrong RD Delhi WI 53543","LIB_FEAT_ID":26,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772 Armstrong RD Delhi WI 53543","LIB_FEAT_ID":26,"USED_FOR_CAND":"N","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"ADDR_KEY":[{"FEAT_DESC":"772|ARMSTRNK||53543","LIB_FEAT_ID":37,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772|ARMSTRNK||53543","LIB_FEAT_ID":37,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"772|ARMSTRNK||TL","LIB_FEAT_ID":17,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772|ARMSTRNK||TL","LIB_FEAT_ID":17,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"DOB":[{"FEAT_DESC":"6/9/1983","LIB_FEAT_ID":25,"FEAT_DESC_VALUES":[{"FEAT_DESC":"6/9/1983","LIB_FEAT_ID":25,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"GENDER":[{"FEAT_DESC":"F","LIB_FEAT_ID":3,"FEAT_DESC_VALUES":[{"FEAT_DESC":"F","LIB_FEAT_ID":3,"USED_FOR_CAND":"N","USED_FOR_SCORING":"Y","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"ID_KEY":[{"FEAT_DESC":"ACCT_NUM=5534202208773608","LIB_FEAT_ID":19,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ACCT_NUM=5534202208773608","LIB_FEAT_ID":19,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"SSN=153-33-5185","LIB_FEAT_ID":38,"FEAT_DESC_VALUES":[{"FEAT_DESC":"SSN=153-33-5185","LIB_FEAT_ID":38,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"LOGIN_ID":[{"FEAT_DESC":"flavorh2","LIB_FEAT_ID":28,"FEAT_DESC_VALUES":[{"FEAT_DESC":"flavorh2","LIB_FEAT_ID":28,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"NAME":[{"FEAT_DESC":"OCEANGUY","LIB_FEAT_ID":24,"FEAT_DESC_VALUES":[{"FEAT_DESC":"OCEANGUY","LIB_FEAT_ID":24,"USED_FOR_CAND":"N","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"NAME_KEY":[{"FEAT_DESC":"ASNK","LIB_FEAT_ID":29,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ASNK","LIB_FEAT_ID":29,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"ASNK|ADDRESS.CITY_STD=TL","LIB_FEAT_ID":34,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ASNK|ADDRESS.CITY_STD=TL","LIB_FEAT_ID":34,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"ASNK|DOB.MMDD_HASH=0906","LIB_FEAT_ID":32,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ASNK|DOB.MMDD_HASH=0906","LIB_FEAT_ID":32,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"ASNK|DOB.MMYY_HASH=0683","LIB_FEAT_ID":30,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ASNK|DOB.MMYY_HASH=0683","LIB_FEAT_ID":30,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"ASNK|DOB=80906","LIB_FEAT_ID":31,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ASNK|DOB=80906","LIB_FEAT_ID":31,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"ASNK|PHONE.PHONE_LAST_5=10796","LIB_FEAT_ID":33,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ASNK|PHONE.PHONE_LAST_5=10796","LIB_FEAT_ID":33,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"ASNK|POST=53543","LIB_FEAT_ID":36,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ASNK|POST=53543","LIB_FEAT_ID":36,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"ASNK|SSN=5185","LIB_FEAT_ID":35,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ASNK|SSN=5185","LIB_FEAT_ID":35,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"PHONE":[{"FEAT_DESC":"225-671-0796","LIB_FEAT_ID":5,"FEAT_DESC_VALUES":[{"FEAT_DESC":"225-671-0796","LIB_FEAT_ID":5,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"PHONE_KEY":[{"FEAT_DESC":"2256710796","LIB_FEAT_ID":21,"FEAT_DESC_VALUES":[{"FEAT_DESC":"2256710796","LIB_FEAT_ID":21,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"SEARCH_KEY":[{"FEAT_DESC":"LOGIN_ID:FLAVORH2|","LIB_FEAT_ID":40,"FEAT_DESC_VALUES":[{"FEAT_DESC":"LOGIN_ID:FLAVORH2|","LIB_FEAT_ID":40,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"SSN:5185|80906|","LIB_FEAT_ID":39,"FEAT_DESC_VALUES":[{"FEAT_DESC":"SSN:5185|80906|","LIB_FEAT_ID":39,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"SSN":[{"FEAT_DESC":"153-33-5185","LIB_FEAT_ID":27,"FEAT_DESC_VALUES":[{"FEAT_DESC":"153-33-5185","LIB_FEAT_ID":27,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}]},"RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-12-06 16:13:27.208","LAST_SEEN_DT":"2022-12-06 16:13:27.208"}],"LAST_SEEN_DT":"2022-12-06 16:13:27.208","RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":"222","ENTITY_TYPE":"TEST","INTERNAL_ID":2,"ENTITY_KEY":"740BA22D15CA88462A930AF8A7C904FF5E48226C","ENTITY_DESC":"OCEANGUY","MATCH_KEY":"","MATCH_LEVEL":0,"MATCH_LEVEL_CODE":"","ERRULE_CODE":"","LAST_SEEN_DT":"2022-12-06 16:13:27.208","FEATURES":[{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":24},{"LIB_FEAT_ID":25},{"LIB_FEAT_ID":26},{"LIB_FEAT_ID":27},{"LIB_FEAT_ID":28},{"LIB_FEAT_ID":29},{"LIB_FEAT_ID":30},{"LIB_FEAT_ID":31},{"LIB_FEAT_ID":32},{"LIB_FEAT_ID":33},{"LIB_FEAT_ID":34},{"LIB_FEAT_ID":35},{"LIB_FEAT_ID":36},{"LIB_FEAT_ID":37},{"LIB_FEAT_ID":38},{"LIB_FEAT_ID":39},{"LIB_FEAT_ID":40}]}]},"RELATED_ENTITIES":[{"ENTITY_ID":1,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0,"ENTITY_NAME":"JOHNSON","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":6,"FIRST_SEEN_DT":"2022-12-06 16:13:27.135","LAST_SEEN_DT":"2022-12-06 16:13:27.916"}],"LAST_SEEN_DT":"2022-12-06 16:13:27.916"},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+ADDRESS+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0,"ENTITY_NAME":"Smith","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-12-06 16:13:27.272","LAST_SEEN_DT":"2022-12-06 16:13:27.272"}],"LAST_SEEN_DT":"2022-12-06 16:13:27.272"}]}]}`
*/
func (client *G2engine) WhyRecords(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string) (string, error) {
	if client.isTrace {
		client.traceEntry(153, dataSourceCode1, recordID1, dataSourceCode2, recordID2)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"dataSourceCode1": dataSourceCode1,
				"recordID1":       recordID1,
				"dataSourceCode2": dataSourceCode2,
				"recordID2":       recordID2,
			}
			client.notify(ctx, 8073, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(154, dataSourceCode1, recordID1, dataSourceCode2, recordID2, client.WhyRecordsResult, err, time.Since(entryTime))
	}
	return client.WhyRecordsResult, err
}

/*
The WhyRecords_V2 method explains why records belong to their resolved entities.
It extends WhyRecords() by adding output control flags.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode1: Identifies the provenance of the data.
  - recordID1: The unique identifier within the records of the same data source.
  - dataSourceCode2: Identifies the provenance of the data.
  - recordID2: The unique identifier within the records of the same data source.
  - flags: Flags used to control information returned.

Output
  - A JSON document.
    See the example output.
*/
func (client *G2engine) WhyRecords_V2(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, flags int64) (string, error) {
	if client.isTrace {
		client.traceEntry(155, dataSourceCode1, recordID1, dataSourceCode2, recordID2, flags)
	}
	var err error = nil
	entryTime := time.Now()
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"dataSourceCode1": dataSourceCode1,
				"recordID1":       recordID1,
				"dataSourceCode2": dataSourceCode2,
				"recordID2":       recordID2,
			}
			client.notify(ctx, 8074, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(156, dataSourceCode1, recordID1, dataSourceCode2, recordID2, flags, client.WhyRecords_V2Result, err, time.Since(entryTime))
	}
	return client.WhyRecords_V2Result, err
}
