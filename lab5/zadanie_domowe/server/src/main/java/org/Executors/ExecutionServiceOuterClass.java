// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: ExecutionService.proto

package org.Executors;

public final class ExecutionServiceOuterClass {
  private ExecutionServiceOuterClass() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_org_Executors_ExecutionRequest_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_org_Executors_ExecutionRequest_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_org_Executors_ExecutionResponse_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_org_Executors_ExecutionResponse_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\026ExecutionService.proto\022\rorg.Executors\"" +
      "\\\n\020ExecutionRequest\022\021\n\tclassName\030\001 \001(\t\022\022" +
      "\n\nmethodName\030\002 \001(\t\022\023\n\013jarLocation\030\003 \001(\t\022" +
      "\014\n\004data\030\004 \001(\t\"2\n\021ExecutionResponse\022\017\n\007er" +
      "rCode\030\001 \001(\t\022\014\n\004data\030\002 \001(\t2`\n\020ExecutionSe" +
      "rvice\022L\n\007execute\022\037.org.Executors.Executi" +
      "onRequest\032 .org.Executors.ExecutionRespo" +
      "nseB\026P\001Z\022/dynamic_executorsb\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
        });
    internal_static_org_Executors_ExecutionRequest_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_org_Executors_ExecutionRequest_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_org_Executors_ExecutionRequest_descriptor,
        new java.lang.String[] { "ClassName", "MethodName", "JarLocation", "Data", });
    internal_static_org_Executors_ExecutionResponse_descriptor =
      getDescriptor().getMessageTypes().get(1);
    internal_static_org_Executors_ExecutionResponse_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_org_Executors_ExecutionResponse_descriptor,
        new java.lang.String[] { "ErrCode", "Data", });
  }

  // @@protoc_insertion_point(outer_class_scope)
}
