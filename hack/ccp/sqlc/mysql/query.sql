-- name: CreateJob :exec
INSERT INTO Jobs (jobUUID, jobType, createdTime, createdTimeDec, directory, SIPUUID, unitType, currentStep, microserviceGroup, hidden, MicroServiceChainLinksPK, subJobOf) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: UpdateJobStatus :exec
UPDATE Jobs SET currentStep = ? WHERE Jobs.jobUUID = ?;

-- name: CleanUpTasksWithAwaitingJobs :exec
DELETE FROM Tasks WHERE jobuuid IN (SELECT jobUUID FROM Jobs WHERE currentStep = 1);

-- name: CleanUpAwaitingJobs :exec
DELETE FROM Jobs WHERE currentStep = 1;

-- name: CleanUpActiveJobs :exec
UPDATE Jobs SET currentStep = 4 WHERE currentStep = 3;

-- name: CleanUpActiveTransfers :exec
UPDATE Transfers SET status = 4, completed_at = UTC_TIMESTAMP() WHERE status IN (0, 1);

-- name: CleanUpActiveSIPs :exec
UPDATE SIPs SET status = 4, completed_at = UTC_TIMESTAMP() WHERE status IN (0, 1);

-- name: CleanUpActiveTasks :exec
UPDATE Tasks SET exitCode = -1, stdError = "MCP shut down while processing." WHERE exitCode IS NULL;

-- name: ReadWorkflowUnitVariable :one
SELECT microServiceChainLink FROM UnitVariables WHERE unitType = ? AND unitUUID = ? AND variable = ?;

-- name: CreateWorkflowUnitVariable :exec
INSERT INTO UnitVariables (unitType, unitUUID, variable, variableValue, microServiceChainLink) VALUES (?, ?, ?, ?, ?)
ON DUPLICATE KEY UPDATE variableValue = VALUES(variableValue), microServiceChainLink = VALUES(microServiceChainLink);

-- name: GetLock :one
SELECT COALESCE(GET_LOCK('lock', 0), 0);

-- name: ReleaseLock :one
SELECT RELEASE_LOCK('lock');
