// Generated by the protocol buffer compiler.  DO NOT EDIT!
// NO CHECKED-IN PROTOBUF GENCODE
// source: student.proto
// Protobuf Java Version: 4.28.0

package org.wso2.apk.integration.utils.clients.student_service;

public final class Student {
  private Student() {
  }

  static {
    com.google.protobuf.RuntimeVersion.validateProtobufGencodeVersion(
        com.google.protobuf.RuntimeVersion.RuntimeDomain.PUBLIC,
        /* major= */ 4,
        /* minor= */ 28,
        /* patch= */ 0,
        /* suffix= */ "",
        Student.class.getName());
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }

  static final com.google.protobuf.Descriptors.Descriptor internal_static_org_apk_v1_student_service_StudentRequest_descriptor;
  static final com.google.protobuf.GeneratedMessage.FieldAccessorTable internal_static_org_apk_v1_student_service_StudentRequest_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor internal_static_org_apk_v1_student_service_StudentResponse_descriptor;
  static final com.google.protobuf.GeneratedMessage.FieldAccessorTable internal_static_org_apk_v1_student_service_StudentResponse_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor getDescriptor() {
    return descriptor;
  }

  private static com.google.protobuf.Descriptors.FileDescriptor descriptor;
  static {
    java.lang.String[] descriptorData = {
        "\n\rstudent.proto\022\032org.apk.v1.student_serv" +
            "ice\"\034\n\016StudentRequest\022\n\n\002id\030\003 \001(\005\",\n\017Stu" +
            "dentResponse\022\014\n\004name\030\001 \001(\t\022\013\n\003age\030\002 \001(\0052" +
            "\326\003\n\016StudentService\022g\n\nGetStudent\022*.org.a" +
            "pk.v1.student_service.StudentRequest\032+.o" +
            "rg.apk.v1.student_service.StudentRespons" +
            "e\"\000\022o\n\020GetStudentStream\022*.org.apk.v1.stu" +
            "dent_service.StudentRequest\032+.org.apk.v1" +
            ".student_service.StudentResponse\"\0000\001\022p\n\021" +
            "SendStudentStream\022*.org.apk.v1.student_s" +
            "ervice.StudentRequest\032+.org.apk.v1.stude" +
            "nt_service.StudentResponse\"\000(\001\022x\n\027SendAn" +
            "dGetStudentStream\022*.org.apk.v1.student_s" +
            "ervice.StudentRequest\032+.org.apk.v1.stude" +
            "nt_service.StudentResponse\"\000(\0010\001B\036\n\032org." +
            "apk.v1.student_serviceP\001b\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
        .internalBuildGeneratedFileFrom(descriptorData,
            new com.google.protobuf.Descriptors.FileDescriptor[] {
            });
    internal_static_org_apk_v1_student_service_StudentRequest_descriptor = getDescriptor().getMessageTypes().get(0);
    internal_static_org_apk_v1_student_service_StudentRequest_fieldAccessorTable = new com.google.protobuf.GeneratedMessage.FieldAccessorTable(
        internal_static_org_apk_v1_student_service_StudentRequest_descriptor,
        new java.lang.String[] { "Id", });
    internal_static_org_apk_v1_student_service_StudentResponse_descriptor = getDescriptor().getMessageTypes().get(1);
    internal_static_org_apk_v1_student_service_StudentResponse_fieldAccessorTable = new com.google.protobuf.GeneratedMessage.FieldAccessorTable(
        internal_static_org_apk_v1_student_service_StudentResponse_descriptor,
        new java.lang.String[] { "Name", "Age", });
    descriptor.resolveAllFeaturesImmutable();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
