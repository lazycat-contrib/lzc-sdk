// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: localdevice/device.proto

package cloud.lazycat.apis.localdevice;

public final class Device {
  private Device() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  public interface DeviceInfoOrBuilder extends
      // @@protoc_insertion_point(interface_extends:cloud.lazycat.apis.localdevice.DeviceInfo)
      com.google.protobuf.MessageOrBuilder {

    /**
     * <code>string OS = 1;</code>
     * @return The oS.
     */
    java.lang.String getOS();
    /**
     * <code>string OS = 1;</code>
     * @return The bytes for oS.
     */
    com.google.protobuf.ByteString
        getOSBytes();

    /**
     * <code>string CPU = 2;</code>
     * @return The cPU.
     */
    java.lang.String getCPU();
    /**
     * <code>string CPU = 2;</code>
     * @return The bytes for cPU.
     */
    com.google.protobuf.ByteString
        getCPUBytes();

    /**
     * <code>string name = 3;</code>
     * @return The name.
     */
    java.lang.String getName();
    /**
     * <code>string name = 3;</code>
     * @return The bytes for name.
     */
    com.google.protobuf.ByteString
        getNameBytes();

    /**
     * <code>string documentRootDir = 4;</code>
     * @return The documentRootDir.
     */
    java.lang.String getDocumentRootDir();
    /**
     * <code>string documentRootDir = 4;</code>
     * @return The bytes for documentRootDir.
     */
    com.google.protobuf.ByteString
        getDocumentRootDirBytes();
  }
  /**
   * Protobuf type {@code cloud.lazycat.apis.localdevice.DeviceInfo}
   */
  public static final class DeviceInfo extends
      com.google.protobuf.GeneratedMessageV3 implements
      // @@protoc_insertion_point(message_implements:cloud.lazycat.apis.localdevice.DeviceInfo)
      DeviceInfoOrBuilder {
  private static final long serialVersionUID = 0L;
    // Use DeviceInfo.newBuilder() to construct.
    private DeviceInfo(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
      super(builder);
    }
    private DeviceInfo() {
      oS_ = "";
      cPU_ = "";
      name_ = "";
      documentRootDir_ = "";
    }

    @java.lang.Override
    @SuppressWarnings({"unused"})
    protected java.lang.Object newInstance(
        UnusedPrivateParameter unused) {
      return new DeviceInfo();
    }

    @java.lang.Override
    public final com.google.protobuf.UnknownFieldSet
    getUnknownFields() {
      return this.unknownFields;
    }
    private DeviceInfo(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      this();
      if (extensionRegistry == null) {
        throw new java.lang.NullPointerException();
      }
      com.google.protobuf.UnknownFieldSet.Builder unknownFields =
          com.google.protobuf.UnknownFieldSet.newBuilder();
      try {
        boolean done = false;
        while (!done) {
          int tag = input.readTag();
          switch (tag) {
            case 0:
              done = true;
              break;
            case 10: {
              java.lang.String s = input.readStringRequireUtf8();

              oS_ = s;
              break;
            }
            case 18: {
              java.lang.String s = input.readStringRequireUtf8();

              cPU_ = s;
              break;
            }
            case 26: {
              java.lang.String s = input.readStringRequireUtf8();

              name_ = s;
              break;
            }
            case 34: {
              java.lang.String s = input.readStringRequireUtf8();

              documentRootDir_ = s;
              break;
            }
            default: {
              if (!parseUnknownField(
                  input, unknownFields, extensionRegistry, tag)) {
                done = true;
              }
              break;
            }
          }
        }
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        throw e.setUnfinishedMessage(this);
      } catch (com.google.protobuf.UninitializedMessageException e) {
        throw e.asInvalidProtocolBufferException().setUnfinishedMessage(this);
      } catch (java.io.IOException e) {
        throw new com.google.protobuf.InvalidProtocolBufferException(
            e).setUnfinishedMessage(this);
      } finally {
        this.unknownFields = unknownFields.build();
        makeExtensionsImmutable();
      }
    }
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return cloud.lazycat.apis.localdevice.Device.internal_static_cloud_lazycat_apis_localdevice_DeviceInfo_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return cloud.lazycat.apis.localdevice.Device.internal_static_cloud_lazycat_apis_localdevice_DeviceInfo_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              cloud.lazycat.apis.localdevice.Device.DeviceInfo.class, cloud.lazycat.apis.localdevice.Device.DeviceInfo.Builder.class);
    }

    public static final int OS_FIELD_NUMBER = 1;
    private volatile java.lang.Object oS_;
    /**
     * <code>string OS = 1;</code>
     * @return The oS.
     */
    @java.lang.Override
    public java.lang.String getOS() {
      java.lang.Object ref = oS_;
      if (ref instanceof java.lang.String) {
        return (java.lang.String) ref;
      } else {
        com.google.protobuf.ByteString bs = 
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        oS_ = s;
        return s;
      }
    }
    /**
     * <code>string OS = 1;</code>
     * @return The bytes for oS.
     */
    @java.lang.Override
    public com.google.protobuf.ByteString
        getOSBytes() {
      java.lang.Object ref = oS_;
      if (ref instanceof java.lang.String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        oS_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }

    public static final int CPU_FIELD_NUMBER = 2;
    private volatile java.lang.Object cPU_;
    /**
     * <code>string CPU = 2;</code>
     * @return The cPU.
     */
    @java.lang.Override
    public java.lang.String getCPU() {
      java.lang.Object ref = cPU_;
      if (ref instanceof java.lang.String) {
        return (java.lang.String) ref;
      } else {
        com.google.protobuf.ByteString bs = 
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        cPU_ = s;
        return s;
      }
    }
    /**
     * <code>string CPU = 2;</code>
     * @return The bytes for cPU.
     */
    @java.lang.Override
    public com.google.protobuf.ByteString
        getCPUBytes() {
      java.lang.Object ref = cPU_;
      if (ref instanceof java.lang.String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        cPU_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }

    public static final int NAME_FIELD_NUMBER = 3;
    private volatile java.lang.Object name_;
    /**
     * <code>string name = 3;</code>
     * @return The name.
     */
    @java.lang.Override
    public java.lang.String getName() {
      java.lang.Object ref = name_;
      if (ref instanceof java.lang.String) {
        return (java.lang.String) ref;
      } else {
        com.google.protobuf.ByteString bs = 
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        name_ = s;
        return s;
      }
    }
    /**
     * <code>string name = 3;</code>
     * @return The bytes for name.
     */
    @java.lang.Override
    public com.google.protobuf.ByteString
        getNameBytes() {
      java.lang.Object ref = name_;
      if (ref instanceof java.lang.String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        name_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }

    public static final int DOCUMENTROOTDIR_FIELD_NUMBER = 4;
    private volatile java.lang.Object documentRootDir_;
    /**
     * <code>string documentRootDir = 4;</code>
     * @return The documentRootDir.
     */
    @java.lang.Override
    public java.lang.String getDocumentRootDir() {
      java.lang.Object ref = documentRootDir_;
      if (ref instanceof java.lang.String) {
        return (java.lang.String) ref;
      } else {
        com.google.protobuf.ByteString bs = 
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        documentRootDir_ = s;
        return s;
      }
    }
    /**
     * <code>string documentRootDir = 4;</code>
     * @return The bytes for documentRootDir.
     */
    @java.lang.Override
    public com.google.protobuf.ByteString
        getDocumentRootDirBytes() {
      java.lang.Object ref = documentRootDir_;
      if (ref instanceof java.lang.String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        documentRootDir_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }

    private byte memoizedIsInitialized = -1;
    @java.lang.Override
    public final boolean isInitialized() {
      byte isInitialized = memoizedIsInitialized;
      if (isInitialized == 1) return true;
      if (isInitialized == 0) return false;

      memoizedIsInitialized = 1;
      return true;
    }

    @java.lang.Override
    public void writeTo(com.google.protobuf.CodedOutputStream output)
                        throws java.io.IOException {
      if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(oS_)) {
        com.google.protobuf.GeneratedMessageV3.writeString(output, 1, oS_);
      }
      if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(cPU_)) {
        com.google.protobuf.GeneratedMessageV3.writeString(output, 2, cPU_);
      }
      if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(name_)) {
        com.google.protobuf.GeneratedMessageV3.writeString(output, 3, name_);
      }
      if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(documentRootDir_)) {
        com.google.protobuf.GeneratedMessageV3.writeString(output, 4, documentRootDir_);
      }
      unknownFields.writeTo(output);
    }

    @java.lang.Override
    public int getSerializedSize() {
      int size = memoizedSize;
      if (size != -1) return size;

      size = 0;
      if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(oS_)) {
        size += com.google.protobuf.GeneratedMessageV3.computeStringSize(1, oS_);
      }
      if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(cPU_)) {
        size += com.google.protobuf.GeneratedMessageV3.computeStringSize(2, cPU_);
      }
      if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(name_)) {
        size += com.google.protobuf.GeneratedMessageV3.computeStringSize(3, name_);
      }
      if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(documentRootDir_)) {
        size += com.google.protobuf.GeneratedMessageV3.computeStringSize(4, documentRootDir_);
      }
      size += unknownFields.getSerializedSize();
      memoizedSize = size;
      return size;
    }

    @java.lang.Override
    public boolean equals(final java.lang.Object obj) {
      if (obj == this) {
       return true;
      }
      if (!(obj instanceof cloud.lazycat.apis.localdevice.Device.DeviceInfo)) {
        return super.equals(obj);
      }
      cloud.lazycat.apis.localdevice.Device.DeviceInfo other = (cloud.lazycat.apis.localdevice.Device.DeviceInfo) obj;

      if (!getOS()
          .equals(other.getOS())) return false;
      if (!getCPU()
          .equals(other.getCPU())) return false;
      if (!getName()
          .equals(other.getName())) return false;
      if (!getDocumentRootDir()
          .equals(other.getDocumentRootDir())) return false;
      if (!unknownFields.equals(other.unknownFields)) return false;
      return true;
    }

    @java.lang.Override
    public int hashCode() {
      if (memoizedHashCode != 0) {
        return memoizedHashCode;
      }
      int hash = 41;
      hash = (19 * hash) + getDescriptor().hashCode();
      hash = (37 * hash) + OS_FIELD_NUMBER;
      hash = (53 * hash) + getOS().hashCode();
      hash = (37 * hash) + CPU_FIELD_NUMBER;
      hash = (53 * hash) + getCPU().hashCode();
      hash = (37 * hash) + NAME_FIELD_NUMBER;
      hash = (53 * hash) + getName().hashCode();
      hash = (37 * hash) + DOCUMENTROOTDIR_FIELD_NUMBER;
      hash = (53 * hash) + getDocumentRootDir().hashCode();
      hash = (29 * hash) + unknownFields.hashCode();
      memoizedHashCode = hash;
      return hash;
    }

    public static cloud.lazycat.apis.localdevice.Device.DeviceInfo parseFrom(
        java.nio.ByteBuffer data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static cloud.lazycat.apis.localdevice.Device.DeviceInfo parseFrom(
        java.nio.ByteBuffer data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static cloud.lazycat.apis.localdevice.Device.DeviceInfo parseFrom(
        com.google.protobuf.ByteString data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static cloud.lazycat.apis.localdevice.Device.DeviceInfo parseFrom(
        com.google.protobuf.ByteString data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static cloud.lazycat.apis.localdevice.Device.DeviceInfo parseFrom(byte[] data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static cloud.lazycat.apis.localdevice.Device.DeviceInfo parseFrom(
        byte[] data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static cloud.lazycat.apis.localdevice.Device.DeviceInfo parseFrom(java.io.InputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input);
    }
    public static cloud.lazycat.apis.localdevice.Device.DeviceInfo parseFrom(
        java.io.InputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input, extensionRegistry);
    }
    public static cloud.lazycat.apis.localdevice.Device.DeviceInfo parseDelimitedFrom(java.io.InputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseDelimitedWithIOException(PARSER, input);
    }
    public static cloud.lazycat.apis.localdevice.Device.DeviceInfo parseDelimitedFrom(
        java.io.InputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
    }
    public static cloud.lazycat.apis.localdevice.Device.DeviceInfo parseFrom(
        com.google.protobuf.CodedInputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input);
    }
    public static cloud.lazycat.apis.localdevice.Device.DeviceInfo parseFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input, extensionRegistry);
    }

    @java.lang.Override
    public Builder newBuilderForType() { return newBuilder(); }
    public static Builder newBuilder() {
      return DEFAULT_INSTANCE.toBuilder();
    }
    public static Builder newBuilder(cloud.lazycat.apis.localdevice.Device.DeviceInfo prototype) {
      return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
    }
    @java.lang.Override
    public Builder toBuilder() {
      return this == DEFAULT_INSTANCE
          ? new Builder() : new Builder().mergeFrom(this);
    }

    @java.lang.Override
    protected Builder newBuilderForType(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      Builder builder = new Builder(parent);
      return builder;
    }
    /**
     * Protobuf type {@code cloud.lazycat.apis.localdevice.DeviceInfo}
     */
    public static final class Builder extends
        com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
        // @@protoc_insertion_point(builder_implements:cloud.lazycat.apis.localdevice.DeviceInfo)
        cloud.lazycat.apis.localdevice.Device.DeviceInfoOrBuilder {
      public static final com.google.protobuf.Descriptors.Descriptor
          getDescriptor() {
        return cloud.lazycat.apis.localdevice.Device.internal_static_cloud_lazycat_apis_localdevice_DeviceInfo_descriptor;
      }

      @java.lang.Override
      protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
          internalGetFieldAccessorTable() {
        return cloud.lazycat.apis.localdevice.Device.internal_static_cloud_lazycat_apis_localdevice_DeviceInfo_fieldAccessorTable
            .ensureFieldAccessorsInitialized(
                cloud.lazycat.apis.localdevice.Device.DeviceInfo.class, cloud.lazycat.apis.localdevice.Device.DeviceInfo.Builder.class);
      }

      // Construct using cloud.lazycat.apis.localdevice.Device.DeviceInfo.newBuilder()
      private Builder() {
        maybeForceBuilderInitialization();
      }

      private Builder(
          com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
        super(parent);
        maybeForceBuilderInitialization();
      }
      private void maybeForceBuilderInitialization() {
        if (com.google.protobuf.GeneratedMessageV3
                .alwaysUseFieldBuilders) {
        }
      }
      @java.lang.Override
      public Builder clear() {
        super.clear();
        oS_ = "";

        cPU_ = "";

        name_ = "";

        documentRootDir_ = "";

        return this;
      }

      @java.lang.Override
      public com.google.protobuf.Descriptors.Descriptor
          getDescriptorForType() {
        return cloud.lazycat.apis.localdevice.Device.internal_static_cloud_lazycat_apis_localdevice_DeviceInfo_descriptor;
      }

      @java.lang.Override
      public cloud.lazycat.apis.localdevice.Device.DeviceInfo getDefaultInstanceForType() {
        return cloud.lazycat.apis.localdevice.Device.DeviceInfo.getDefaultInstance();
      }

      @java.lang.Override
      public cloud.lazycat.apis.localdevice.Device.DeviceInfo build() {
        cloud.lazycat.apis.localdevice.Device.DeviceInfo result = buildPartial();
        if (!result.isInitialized()) {
          throw newUninitializedMessageException(result);
        }
        return result;
      }

      @java.lang.Override
      public cloud.lazycat.apis.localdevice.Device.DeviceInfo buildPartial() {
        cloud.lazycat.apis.localdevice.Device.DeviceInfo result = new cloud.lazycat.apis.localdevice.Device.DeviceInfo(this);
        result.oS_ = oS_;
        result.cPU_ = cPU_;
        result.name_ = name_;
        result.documentRootDir_ = documentRootDir_;
        onBuilt();
        return result;
      }

      @java.lang.Override
      public Builder clone() {
        return super.clone();
      }
      @java.lang.Override
      public Builder setField(
          com.google.protobuf.Descriptors.FieldDescriptor field,
          java.lang.Object value) {
        return super.setField(field, value);
      }
      @java.lang.Override
      public Builder clearField(
          com.google.protobuf.Descriptors.FieldDescriptor field) {
        return super.clearField(field);
      }
      @java.lang.Override
      public Builder clearOneof(
          com.google.protobuf.Descriptors.OneofDescriptor oneof) {
        return super.clearOneof(oneof);
      }
      @java.lang.Override
      public Builder setRepeatedField(
          com.google.protobuf.Descriptors.FieldDescriptor field,
          int index, java.lang.Object value) {
        return super.setRepeatedField(field, index, value);
      }
      @java.lang.Override
      public Builder addRepeatedField(
          com.google.protobuf.Descriptors.FieldDescriptor field,
          java.lang.Object value) {
        return super.addRepeatedField(field, value);
      }
      @java.lang.Override
      public Builder mergeFrom(com.google.protobuf.Message other) {
        if (other instanceof cloud.lazycat.apis.localdevice.Device.DeviceInfo) {
          return mergeFrom((cloud.lazycat.apis.localdevice.Device.DeviceInfo)other);
        } else {
          super.mergeFrom(other);
          return this;
        }
      }

      public Builder mergeFrom(cloud.lazycat.apis.localdevice.Device.DeviceInfo other) {
        if (other == cloud.lazycat.apis.localdevice.Device.DeviceInfo.getDefaultInstance()) return this;
        if (!other.getOS().isEmpty()) {
          oS_ = other.oS_;
          onChanged();
        }
        if (!other.getCPU().isEmpty()) {
          cPU_ = other.cPU_;
          onChanged();
        }
        if (!other.getName().isEmpty()) {
          name_ = other.name_;
          onChanged();
        }
        if (!other.getDocumentRootDir().isEmpty()) {
          documentRootDir_ = other.documentRootDir_;
          onChanged();
        }
        this.mergeUnknownFields(other.unknownFields);
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
        cloud.lazycat.apis.localdevice.Device.DeviceInfo parsedMessage = null;
        try {
          parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
        } catch (com.google.protobuf.InvalidProtocolBufferException e) {
          parsedMessage = (cloud.lazycat.apis.localdevice.Device.DeviceInfo) e.getUnfinishedMessage();
          throw e.unwrapIOException();
        } finally {
          if (parsedMessage != null) {
            mergeFrom(parsedMessage);
          }
        }
        return this;
      }

      private java.lang.Object oS_ = "";
      /**
       * <code>string OS = 1;</code>
       * @return The oS.
       */
      public java.lang.String getOS() {
        java.lang.Object ref = oS_;
        if (!(ref instanceof java.lang.String)) {
          com.google.protobuf.ByteString bs =
              (com.google.protobuf.ByteString) ref;
          java.lang.String s = bs.toStringUtf8();
          oS_ = s;
          return s;
        } else {
          return (java.lang.String) ref;
        }
      }
      /**
       * <code>string OS = 1;</code>
       * @return The bytes for oS.
       */
      public com.google.protobuf.ByteString
          getOSBytes() {
        java.lang.Object ref = oS_;
        if (ref instanceof String) {
          com.google.protobuf.ByteString b = 
              com.google.protobuf.ByteString.copyFromUtf8(
                  (java.lang.String) ref);
          oS_ = b;
          return b;
        } else {
          return (com.google.protobuf.ByteString) ref;
        }
      }
      /**
       * <code>string OS = 1;</code>
       * @param value The oS to set.
       * @return This builder for chaining.
       */
      public Builder setOS(
          java.lang.String value) {
        if (value == null) {
    throw new NullPointerException();
  }
  
        oS_ = value;
        onChanged();
        return this;
      }
      /**
       * <code>string OS = 1;</code>
       * @return This builder for chaining.
       */
      public Builder clearOS() {
        
        oS_ = getDefaultInstance().getOS();
        onChanged();
        return this;
      }
      /**
       * <code>string OS = 1;</code>
       * @param value The bytes for oS to set.
       * @return This builder for chaining.
       */
      public Builder setOSBytes(
          com.google.protobuf.ByteString value) {
        if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
        
        oS_ = value;
        onChanged();
        return this;
      }

      private java.lang.Object cPU_ = "";
      /**
       * <code>string CPU = 2;</code>
       * @return The cPU.
       */
      public java.lang.String getCPU() {
        java.lang.Object ref = cPU_;
        if (!(ref instanceof java.lang.String)) {
          com.google.protobuf.ByteString bs =
              (com.google.protobuf.ByteString) ref;
          java.lang.String s = bs.toStringUtf8();
          cPU_ = s;
          return s;
        } else {
          return (java.lang.String) ref;
        }
      }
      /**
       * <code>string CPU = 2;</code>
       * @return The bytes for cPU.
       */
      public com.google.protobuf.ByteString
          getCPUBytes() {
        java.lang.Object ref = cPU_;
        if (ref instanceof String) {
          com.google.protobuf.ByteString b = 
              com.google.protobuf.ByteString.copyFromUtf8(
                  (java.lang.String) ref);
          cPU_ = b;
          return b;
        } else {
          return (com.google.protobuf.ByteString) ref;
        }
      }
      /**
       * <code>string CPU = 2;</code>
       * @param value The cPU to set.
       * @return This builder for chaining.
       */
      public Builder setCPU(
          java.lang.String value) {
        if (value == null) {
    throw new NullPointerException();
  }
  
        cPU_ = value;
        onChanged();
        return this;
      }
      /**
       * <code>string CPU = 2;</code>
       * @return This builder for chaining.
       */
      public Builder clearCPU() {
        
        cPU_ = getDefaultInstance().getCPU();
        onChanged();
        return this;
      }
      /**
       * <code>string CPU = 2;</code>
       * @param value The bytes for cPU to set.
       * @return This builder for chaining.
       */
      public Builder setCPUBytes(
          com.google.protobuf.ByteString value) {
        if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
        
        cPU_ = value;
        onChanged();
        return this;
      }

      private java.lang.Object name_ = "";
      /**
       * <code>string name = 3;</code>
       * @return The name.
       */
      public java.lang.String getName() {
        java.lang.Object ref = name_;
        if (!(ref instanceof java.lang.String)) {
          com.google.protobuf.ByteString bs =
              (com.google.protobuf.ByteString) ref;
          java.lang.String s = bs.toStringUtf8();
          name_ = s;
          return s;
        } else {
          return (java.lang.String) ref;
        }
      }
      /**
       * <code>string name = 3;</code>
       * @return The bytes for name.
       */
      public com.google.protobuf.ByteString
          getNameBytes() {
        java.lang.Object ref = name_;
        if (ref instanceof String) {
          com.google.protobuf.ByteString b = 
              com.google.protobuf.ByteString.copyFromUtf8(
                  (java.lang.String) ref);
          name_ = b;
          return b;
        } else {
          return (com.google.protobuf.ByteString) ref;
        }
      }
      /**
       * <code>string name = 3;</code>
       * @param value The name to set.
       * @return This builder for chaining.
       */
      public Builder setName(
          java.lang.String value) {
        if (value == null) {
    throw new NullPointerException();
  }
  
        name_ = value;
        onChanged();
        return this;
      }
      /**
       * <code>string name = 3;</code>
       * @return This builder for chaining.
       */
      public Builder clearName() {
        
        name_ = getDefaultInstance().getName();
        onChanged();
        return this;
      }
      /**
       * <code>string name = 3;</code>
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
        onChanged();
        return this;
      }

      private java.lang.Object documentRootDir_ = "";
      /**
       * <code>string documentRootDir = 4;</code>
       * @return The documentRootDir.
       */
      public java.lang.String getDocumentRootDir() {
        java.lang.Object ref = documentRootDir_;
        if (!(ref instanceof java.lang.String)) {
          com.google.protobuf.ByteString bs =
              (com.google.protobuf.ByteString) ref;
          java.lang.String s = bs.toStringUtf8();
          documentRootDir_ = s;
          return s;
        } else {
          return (java.lang.String) ref;
        }
      }
      /**
       * <code>string documentRootDir = 4;</code>
       * @return The bytes for documentRootDir.
       */
      public com.google.protobuf.ByteString
          getDocumentRootDirBytes() {
        java.lang.Object ref = documentRootDir_;
        if (ref instanceof String) {
          com.google.protobuf.ByteString b = 
              com.google.protobuf.ByteString.copyFromUtf8(
                  (java.lang.String) ref);
          documentRootDir_ = b;
          return b;
        } else {
          return (com.google.protobuf.ByteString) ref;
        }
      }
      /**
       * <code>string documentRootDir = 4;</code>
       * @param value The documentRootDir to set.
       * @return This builder for chaining.
       */
      public Builder setDocumentRootDir(
          java.lang.String value) {
        if (value == null) {
    throw new NullPointerException();
  }
  
        documentRootDir_ = value;
        onChanged();
        return this;
      }
      /**
       * <code>string documentRootDir = 4;</code>
       * @return This builder for chaining.
       */
      public Builder clearDocumentRootDir() {
        
        documentRootDir_ = getDefaultInstance().getDocumentRootDir();
        onChanged();
        return this;
      }
      /**
       * <code>string documentRootDir = 4;</code>
       * @param value The bytes for documentRootDir to set.
       * @return This builder for chaining.
       */
      public Builder setDocumentRootDirBytes(
          com.google.protobuf.ByteString value) {
        if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
        
        documentRootDir_ = value;
        onChanged();
        return this;
      }
      @java.lang.Override
      public final Builder setUnknownFields(
          final com.google.protobuf.UnknownFieldSet unknownFields) {
        return super.setUnknownFields(unknownFields);
      }

      @java.lang.Override
      public final Builder mergeUnknownFields(
          final com.google.protobuf.UnknownFieldSet unknownFields) {
        return super.mergeUnknownFields(unknownFields);
      }


      // @@protoc_insertion_point(builder_scope:cloud.lazycat.apis.localdevice.DeviceInfo)
    }

    // @@protoc_insertion_point(class_scope:cloud.lazycat.apis.localdevice.DeviceInfo)
    private static final cloud.lazycat.apis.localdevice.Device.DeviceInfo DEFAULT_INSTANCE;
    static {
      DEFAULT_INSTANCE = new cloud.lazycat.apis.localdevice.Device.DeviceInfo();
    }

    public static cloud.lazycat.apis.localdevice.Device.DeviceInfo getDefaultInstance() {
      return DEFAULT_INSTANCE;
    }

    private static final com.google.protobuf.Parser<DeviceInfo>
        PARSER = new com.google.protobuf.AbstractParser<DeviceInfo>() {
      @java.lang.Override
      public DeviceInfo parsePartialFrom(
          com.google.protobuf.CodedInputStream input,
          com.google.protobuf.ExtensionRegistryLite extensionRegistry)
          throws com.google.protobuf.InvalidProtocolBufferException {
        return new DeviceInfo(input, extensionRegistry);
      }
    };

    public static com.google.protobuf.Parser<DeviceInfo> parser() {
      return PARSER;
    }

    @java.lang.Override
    public com.google.protobuf.Parser<DeviceInfo> getParserForType() {
      return PARSER;
    }

    @java.lang.Override
    public cloud.lazycat.apis.localdevice.Device.DeviceInfo getDefaultInstanceForType() {
      return DEFAULT_INSTANCE;
    }

  }

  private static final com.google.protobuf.Descriptors.Descriptor
    internal_static_cloud_lazycat_apis_localdevice_DeviceInfo_descriptor;
  private static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_cloud_lazycat_apis_localdevice_DeviceInfo_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\030localdevice/device.proto\022\036cloud.lazyca" +
      "t.apis.localdevice\032\033google/protobuf/empt" +
      "y.proto\"L\n\nDeviceInfo\022\n\n\002OS\030\001 \001(\t\022\013\n\003CPU" +
      "\030\002 \001(\t\022\014\n\004name\030\003 \001(\t\022\027\n\017documentRootDir\030" +
      "\004 \001(\t2^\n\rDeviceService\022M\n\005Query\022\026.google" +
      ".protobuf.Empty\032*.cloud.lazycat.apis.loc" +
      "aldevice.DeviceInfo\"\000B0Z.gitee.com/linak" +
      "esi/lzc-sdk/lang/go/localdeviceb\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          com.google.protobuf.EmptyProto.getDescriptor(),
        });
    internal_static_cloud_lazycat_apis_localdevice_DeviceInfo_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_cloud_lazycat_apis_localdevice_DeviceInfo_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_cloud_lazycat_apis_localdevice_DeviceInfo_descriptor,
        new java.lang.String[] { "OS", "CPU", "Name", "DocumentRootDir", });
    com.google.protobuf.EmptyProto.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
