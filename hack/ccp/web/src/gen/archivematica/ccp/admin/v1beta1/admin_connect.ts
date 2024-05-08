// @generated by protoc-gen-connect-es v1.4.0 with parameter "target=ts"
// @generated from file archivematica/ccp/admin/v1beta1/admin.proto (package archivematica.ccp.admin.v1beta1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { ApproveTransferRequest, ApproveTransferResponse, CreatePackageRequest, CreatePackageResponse, ListActivePackagesRequest, ListActivePackagesResponse, ListAwaitingDecisionsRequest, ListAwaitingDecisionsResponse, ResolveAwaitingDecisionRequest, ResolveAwaitingDecisionResponse } from "./admin_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service archivematica.ccp.admin.v1beta1.AdminService
 */
export const AdminService = {
  typeName: "archivematica.ccp.admin.v1beta1.AdminService",
  methods: {
    /**
     * @generated from rpc archivematica.ccp.admin.v1beta1.AdminService.CreatePackage
     */
    createPackage: {
      name: "CreatePackage",
      I: CreatePackageRequest,
      O: CreatePackageResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc archivematica.ccp.admin.v1beta1.AdminService.ApproveTransfer
     */
    approveTransfer: {
      name: "ApproveTransfer",
      I: ApproveTransferRequest,
      O: ApproveTransferResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc archivematica.ccp.admin.v1beta1.AdminService.ListActivePackages
     */
    listActivePackages: {
      name: "ListActivePackages",
      I: ListActivePackagesRequest,
      O: ListActivePackagesResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc archivematica.ccp.admin.v1beta1.AdminService.ListAwaitingDecisions
     */
    listAwaitingDecisions: {
      name: "ListAwaitingDecisions",
      I: ListAwaitingDecisionsRequest,
      O: ListAwaitingDecisionsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc archivematica.ccp.admin.v1beta1.AdminService.ResolveAwaitingDecision
     */
    resolveAwaitingDecision: {
      name: "ResolveAwaitingDecision",
      I: ResolveAwaitingDecisionRequest,
      O: ResolveAwaitingDecisionResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

