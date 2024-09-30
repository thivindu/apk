// Generated by the protocol buffer compiler.  DO NOT EDIT!
// NO CHECKED-IN PROTOBUF GENCODE
// source: student.proto
// Protobuf Java Version: 4.28.0

package org.wso2.apk.integration.utils.clients.student_service;

/**
 * Protobuf type {@code org.apk.v1.student_service.StudentResponse}
 */
public final class StudentResponse extends
    com.google.protobuf.GeneratedMessage implements
    // @@protoc_insertion_point(message_implements:org.apk.v1.student_service.StudentResponse)
    StudentResponseOrBuilder {
  private static final long serialVersionUID = 0L;
  static {
    com.google.protobuf.RuntimeVersion.validateProtobufGencodeVersion(
        com.google.protobuf.RuntimeVersion.RuntimeDomain.PUBLIC,
        /* major= */ 4,
        /* minor= */ 28,
        /* patch= */ 0,
        /* suffix= */ "",
        StudentResponse.class.getName());
  }

  // Use StudentResponse.newBuilder() to construct.
  private StudentResponse(com.google.protobuf.GeneratedMessage.Builder<?> builder) {
    super(builder);
  }

  private StudentResponse() {
    name_ = "";
  }

  public static final com.google.protobuf.Descriptors.Descriptor getDescriptor() {
    return org.wso2.apk.integration.utils.clients.student_service.Student.internal_static_org_apk_v1_student_service_StudentResponse_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessage.FieldAccessorTable internalGetFieldAccessorTable() {
    return org.wso2.apk.integration.utils.clients.student_service.Student.internal_static_org_apk_v1_student_service_StudentResponse_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            org.wso2.apk.integration.utils.clients.student_service.StudentResponse.class,
            org.wso2.apk.integration.utils.clients.student_service.StudentResponse.Builder.class);
  }

  public static final int NAME_FIELD_NUMBER = 1;
  @SuppressWarnings("serial")
  private volatile java.lang.Object name_ = "";

  /**
   * <code>string name = 1;</code>
   * 
   * @return The name.
   */
  @java.lang.Override
  public java.lang.String getName() {
    java.lang.Object ref = name_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      name_ = s;
      return s;
    }
  }

  /**
   * <code>string name = 1;</code>
   * 
   * @return The bytes for name.
   */
  @java.lang.Override
  public com.google.protobuf.ByteString getNameBytes() {
    java.lang.Object ref = name_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = com.google.protobuf.ByteString.copyFromUtf8(
          (java.lang.String) ref);
      name_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  public static final int AGE_FIELD_NUMBER = 2;
  private int age_ = 0;

  /**
   * <code>int32 age = 2;</code>
   * 
   * @return The age.
   */
  @java.lang.Override
  public int getAge() {
    return age_;
  }

  private byte memoizedIsInitialized = -1;

  @java.lang.Override
  public final boolean isInitialized() {
    byte isInitialized = memoizedIsInitialized;
    if (isInitialized == 1)
      return true;
    if (isInitialized == 0)
      return false;

    memoizedIsInitialized = 1;
    return true;
  }

  @java.lang.Override
  public void writeTo(com.google.protobuf.CodedOutputStream output)
      throws java.io.IOException {
    if (!com.google.protobuf.GeneratedMessage.isStringEmpty(name_)) {
      com.google.protobuf.GeneratedMessage.writeString(output, 1, name_);
    }
    if (age_ != 0) {
      output.writeInt32(2, age_);
    }
    getUnknownFields().writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1)
      return size;

    size = 0;
    if (!com.google.protobuf.GeneratedMessage.isStringEmpty(name_)) {
      size += com.google.protobuf.GeneratedMessage.computeStringSize(1, name_);
    }
    if (age_ != 0) {
      size += com.google.protobuf.CodedOutputStream
          .computeInt32Size(2, age_);
    }
    size += getUnknownFields().getSerializedSize();
    memoizedSize = size;
    return size;
  }

  @java.lang.Override
  public boolean equals(final java.lang.Object obj) {
    if (obj == this) {
      return true;
    }
    if (!(obj instanceof org.wso2.apk.integration.utils.clients.student_service.StudentResponse)) {
      return super.equals(obj);
    }
    org.wso2.apk.integration.utils.clients.student_service.StudentResponse other = (org.wso2.apk.integration.utils.clients.student_service.StudentResponse) obj;

    if (!getName()
        .equals(other.getName()))
      return false;
    if (getAge() != other.getAge())
      return false;
    if (!getUnknownFields().equals(other.getUnknownFields()))
      return false;
    return true;
  }

  @java.lang.Override
  public int hashCode() {
    if (memoizedHashCode != 0) {
      return memoizedHashCode;
    }
    int hash = 41;
    hash = (19 * hash) + getDescriptor().hashCode();
    hash = (37 * hash) + NAME_FIELD_NUMBER;
    hash = (53 * hash) + getName().hashCode();
    hash = (37 * hash) + AGE_FIELD_NUMBER;
    hash = (53 * hash) + getAge();
    hash = (29 * hash) + getUnknownFields().hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static org.wso2.apk.integration.utils.clients.student_service.StudentResponse parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }

  public static org.wso2.apk.integration.utils.clients.student_service.StudentResponse parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }

  public static org.wso2.apk.integration.utils.clients.student_service.StudentResponse parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }

  public static org.wso2.apk.integration.utils.clients.student_service.StudentResponse parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }

  public static org.wso2.apk.integration.utils.clients.student_service.StudentResponse parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }

  public static org.wso2.apk.integration.utils.clients.student_service.StudentResponse parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }

  public static org.wso2.apk.integration.utils.clients.student_service.StudentResponse parseFrom(
      java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessage
        .parseWithIOException(PARSER, input);
  }

  public static org.wso2.apk.integration.utils.clients.student_service.StudentResponse parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessage
        .parseWithIOException(PARSER, input, extensionRegistry);
  }

  public static org.wso2.apk.integration.utils.clients.student_service.StudentResponse parseDelimitedFrom(
      java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessage
        .parseDelimitedWithIOException(PARSER, input);
  }

  public static org.wso2.apk.integration.utils.clients.student_service.StudentResponse parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessage
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }

  public static org.wso2.apk.integration.utils.clients.student_service.StudentResponse parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessage
        .parseWithIOException(PARSER, input);
  }

  public static org.wso2.apk.integration.utils.clients.student_service.StudentResponse parseFrom(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessage
        .parseWithIOException(PARSER, input, extensionRegistry);
  }

  @java.lang.Override
  public Builder newBuilderForType() {
    return newBuilder();
  }

  public static Builder newBuilder() {
    return DEFAULT_INSTANCE.toBuilder();
  }

  public static Builder newBuilder(org.wso2.apk.integration.utils.clients.student_service.StudentResponse prototype) {
    return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
  }

  @java.lang.Override
  public Builder toBuilder() {
    return this == DEFAULT_INSTANCE
        ? new Builder()
        : new Builder().mergeFrom(this);
  }

  @java.lang.Override
  protected Builder newBuilderForType(
      com.google.protobuf.GeneratedMessage.BuilderParent parent) {
    Builder builder = new Builder(parent);
    return builder;
  }

  /**
   * Protobuf type {@code org.apk.v1.student_service.StudentResponse}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessage.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:org.apk.v1.student_service.StudentResponse)
      org.wso2.apk.integration.utils.clients.student_service.StudentResponseOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor getDescriptor() {
      return org.wso2.apk.integration.utils.clients.student_service.Student.internal_static_org_apk_v1_student_service_StudentResponse_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessage.FieldAccessorTable internalGetFieldAccessorTable() {
      return org.wso2.apk.integration.utils.clients.student_service.Student.internal_static_org_apk_v1_student_service_StudentResponse_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              org.wso2.apk.integration.utils.clients.student_service.StudentResponse.class,
              org.wso2.apk.integration.utils.clients.student_service.StudentResponse.Builder.class);
    }

    // Construct using org.apk.v1.student_service.StudentResponse.newBuilder()
    private Builder() {

    }

    private Builder(
        com.google.protobuf.GeneratedMessage.BuilderParent parent) {
      super(parent);

    }

    @java.lang.Override
    public Builder clear() {
      super.clear();
      bitField0_ = 0;
      name_ = "";
      age_ = 0;
      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor getDescriptorForType() {
      return org.wso2.apk.integration.utils.clients.student_service.Student.internal_static_org_apk_v1_student_service_StudentResponse_descriptor;
    }

    @java.lang.Override
    public org.wso2.apk.integration.utils.clients.student_service.StudentResponse getDefaultInstanceForType() {
      return org.wso2.apk.integration.utils.clients.student_service.StudentResponse.getDefaultInstance();
    }

    @java.lang.Override
    public org.wso2.apk.integration.utils.clients.student_service.StudentResponse build() {
      org.wso2.apk.integration.utils.clients.student_service.StudentResponse result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public org.wso2.apk.integration.utils.clients.student_service.StudentResponse buildPartial() {
      org.wso2.apk.integration.utils.clients.student_service.StudentResponse result = new org.wso2.apk.integration.utils.clients.student_service.StudentResponse(
          this);
      if (bitField0_ != 0) {
        buildPartial0(result);
      }
      onBuilt();
      return result;
    }

    private void buildPartial0(org.wso2.apk.integration.utils.clients.student_service.StudentResponse result) {
      int from_bitField0_ = bitField0_;
      if (((from_bitField0_ & 0x00000001) != 0)) {
        result.name_ = name_;
      }
      if (((from_bitField0_ & 0x00000002) != 0)) {
        result.age_ = age_;
      }
    }

    @java.lang.Override
    public Builder mergeFrom(com.google.protobuf.Message other) {
      if (other instanceof org.wso2.apk.integration.utils.clients.student_service.StudentResponse) {
        return mergeFrom((org.wso2.apk.integration.utils.clients.student_service.StudentResponse) other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(org.wso2.apk.integration.utils.clients.student_service.StudentResponse other) {
      if (other == org.wso2.apk.integration.utils.clients.student_service.StudentResponse.getDefaultInstance())
        return this;
      if (!other.getName().isEmpty()) {
        name_ = other.name_;
        bitField0_ |= 0x00000001;
        onChanged();
      }
      if (other.getAge() != 0) {
        setAge(other.getAge());
      }
      this.mergeUnknownFields(other.getUnknownFields());
      onChanged();
      return this;
    }

    @java.lang.Override
    public final boolean isInitialized() {
      return true;
    }

    @java.lang.Override
    public Builder mergeFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      if (extensionRegistry == null) {
        throw new java.lang.NullPointerException();
      }
      try {
        boolean done = false;
        while (!done) {
          int tag = input.readTag();
          switch (tag) {
            case 0:
              done = true;
              break;
            case 10: {
              name_ = input.readStringRequireUtf8();
              bitField0_ |= 0x00000001;
              break;
            } // case 10
            case 16: {
              age_ = input.readInt32();
              bitField0_ |= 0x00000002;
              break;
            } // case 16
            default: {
              if (!super.parseUnknownField(input, extensionRegistry, tag)) {
                done = true; // was an endgroup tag
              }
              break;
            } // default:
          } // switch (tag)
        } // while (!done)
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        throw e.unwrapIOException();
      } finally {
        onChanged();
      } // finally
      return this;
    }

    private int bitField0_;

    private java.lang.Object name_ = "";

    /**
     * <code>string name = 1;</code>
     * 
     * @return The name.
     */
    public java.lang.String getName() {
      java.lang.Object ref = name_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs = (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        name_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }

    /**
     * <code>string name = 1;</code>
     * 
     * @return The bytes for name.
     */
    public com.google.protobuf.ByteString getNameBytes() {
      java.lang.Object ref = name_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = com.google.protobuf.ByteString.copyFromUtf8(
            (java.lang.String) ref);
        name_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }

    /**
     * <code>string name = 1;</code>
     * 
     * @param value The name to set.
     * @return This builder for chaining.
     */
    public Builder setName(
        java.lang.String value) {
      if (value == null) {
        throw new NullPointerException();
      }
      name_ = value;
      bitField0_ |= 0x00000001;
      onChanged();
      return this;
    }

    /**
     * <code>string name = 1;</code>
     * 
     * @return This builder for chaining.
     */
    public Builder clearName() {
      name_ = getDefaultInstance().getName();
      bitField0_ = (bitField0_ & ~0x00000001);
      onChanged();
      return this;
    }

    /**
     * <code>string name = 1;</code>
     * 
     * @param value The bytes for name to set.
     * @return This builder for chaining.
     */
    public Builder setNameBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) {
        throw new NullPointerException();
      }
      checkByteStringIsUtf8(value);
      name_ = value;
      bitField0_ |= 0x00000001;
      onChanged();
      return this;
    }

    private int age_;

    /**
     * <code>int32 age = 2;</code>
     * 
     * @return The age.
     */
    @java.lang.Override
    public int getAge() {
      return age_;
    }

    /**
     * <code>int32 age = 2;</code>
     * 
     * @param value The age to set.
     * @return This builder for chaining.
     */
    public Builder setAge(int value) {

      age_ = value;
      bitField0_ |= 0x00000002;
      onChanged();
      return this;
    }

    /**
     * <code>int32 age = 2;</code>
     * 
     * @return This builder for chaining.
     */
    public Builder clearAge() {
      bitField0_ = (bitField0_ & ~0x00000002);
      age_ = 0;
      onChanged();
      return this;
    }

    // @@protoc_insertion_point(builder_scope:org.apk.v1.student_service.StudentResponse)
  }

  // @@protoc_insertion_point(class_scope:org.apk.v1.student_service.StudentResponse)
  private static final org.wso2.apk.integration.utils.clients.student_service.StudentResponse DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new org.wso2.apk.integration.utils.clients.student_service.StudentResponse();
  }

  public static org.wso2.apk.integration.utils.clients.student_service.StudentResponse getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<StudentResponse> PARSER = new com.google.protobuf.AbstractParser<StudentResponse>() {
    @java.lang.Override
    public StudentResponse parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      Builder builder = newBuilder();
      try {
        builder.mergeFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        throw e.setUnfinishedMessage(builder.buildPartial());
      } catch (com.google.protobuf.UninitializedMessageException e) {
        throw e.asInvalidProtocolBufferException().setUnfinishedMessage(builder.buildPartial());
      } catch (java.io.IOException e) {
        throw new com.google.protobuf.InvalidProtocolBufferException(e)
            .setUnfinishedMessage(builder.buildPartial());
      }
      return builder.buildPartial();
    }
  };

  public static com.google.protobuf.Parser<StudentResponse> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<StudentResponse> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public org.wso2.apk.integration.utils.clients.student_service.StudentResponse getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}
