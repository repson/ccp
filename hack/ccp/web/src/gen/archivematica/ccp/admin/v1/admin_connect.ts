// @generated by protoc-gen-connect-es v1.4.0 with parameter "target=ts"
// @generated from file archivematica/ccp/admin/v1/admin.proto (package archivematica.ccp.admin.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { ListActivePackagesRequest, ListActivePackagesResponse, ListAwaitingDecisionsRequest, ListAwaitingDecisionsResponse, ResolveAwaitingDecisionRequest, ResolveAwaitingDecisionResponse } from "./admin_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service archivematica.ccp.admin.v1.AdminService
 */
export const AdminService = {
  typeName: "archivematica.ccp.admin.v1.AdminService",
  methods: {
    /**
     * @generated from rpc archivematica.ccp.admin.v1.AdminService.ListActivePackages
     */
    listActivePackages: {
      name: "ListActivePackages",
      I: ListActivePackagesRequest,
      O: ListActivePackagesResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc archivematica.ccp.admin.v1.AdminService.ListAwaitingDecisions
     */
    listAwaitingDecisions: {
      name: "ListAwaitingDecisions",
      I: ListAwaitingDecisionsRequest,
      O: ListAwaitingDecisionsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc archivematica.ccp.admin.v1.AdminService.ResolveAwaitingDecision
     */
    resolveAwaitingDecision: {
      name: "ResolveAwaitingDecision",
      I: ResolveAwaitingDecisionRequest,
      O: ResolveAwaitingDecisionResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;
