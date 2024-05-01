// @generated by protoc-gen-connect-es v1.4.0 with parameter "target=ts"
// @generated from file archivematica/ccp/scheduler/v1/scheduler.proto (package archivematica.ccp.scheduler.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { WorkRequest, WorkResponse } from "./scheduler_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service archivematica.ccp.scheduler.v1.SchedulerService
 */
export const SchedulerService = {
  typeName: "archivematica.ccp.scheduler.v1.SchedulerService",
  methods: {
    /**
     * @generated from rpc archivematica.ccp.scheduler.v1.SchedulerService.Work
     */
    work: {
      name: "Work",
      I: WorkRequest,
      O: WorkResponse,
      kind: MethodKind.BiDiStreaming,
    },
  }
} as const;
