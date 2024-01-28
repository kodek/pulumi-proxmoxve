// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.muehlbachler.pulumi.proxmoxve.Storage;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import io.muehlbachler.pulumi.proxmoxve.Storage.inputs.FileSourceFileArgs;
import io.muehlbachler.pulumi.proxmoxve.Storage.inputs.FileSourceRawArgs;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class FileArgs extends com.pulumi.resources.ResourceArgs {

    public static final FileArgs Empty = new FileArgs();

    /**
     * The content type. If not specified, the content
     * type will be inferred from the file extension. Valid values are:
     * 
     */
    @Import(name="contentType")
    private @Nullable Output<String> contentType;

    /**
     * @return The content type. If not specified, the content
     * type will be inferred from the file extension. Valid values are:
     * 
     */
    public Optional<Output<String>> contentType() {
        return Optional.ofNullable(this.contentType);
    }

    /**
     * The datastore id.
     * 
     */
    @Import(name="datastoreId", required=true)
    private Output<String> datastoreId;

    /**
     * @return The datastore id.
     * 
     */
    public Output<String> datastoreId() {
        return this.datastoreId;
    }

    /**
     * The node name.
     * 
     */
    @Import(name="nodeName", required=true)
    private Output<String> nodeName;

    /**
     * @return The node name.
     * 
     */
    public Output<String> nodeName() {
        return this.nodeName;
    }

    /**
     * Whether to overwrite an existing file (defaults to
     * `true`).
     * 
     */
    @Import(name="overwrite")
    private @Nullable Output<Boolean> overwrite;

    /**
     * @return Whether to overwrite an existing file (defaults to
     * `true`).
     * 
     */
    public Optional<Output<Boolean>> overwrite() {
        return Optional.ofNullable(this.overwrite);
    }

    /**
     * The source file (conflicts with `source_raw`),
     * could be a local file or a URL. If the source file is a URL, the file will
     * be downloaded and stored locally before uploading it to Proxmox VE.
     * 
     */
    @Import(name="sourceFile")
    private @Nullable Output<FileSourceFileArgs> sourceFile;

    /**
     * @return The source file (conflicts with `source_raw`),
     * could be a local file or a URL. If the source file is a URL, the file will
     * be downloaded and stored locally before uploading it to Proxmox VE.
     * 
     */
    public Optional<Output<FileSourceFileArgs>> sourceFile() {
        return Optional.ofNullable(this.sourceFile);
    }

    /**
     * The raw source (conflicts with `source_file`).
     * 
     */
    @Import(name="sourceRaw")
    private @Nullable Output<FileSourceRawArgs> sourceRaw;

    /**
     * @return The raw source (conflicts with `source_file`).
     * 
     */
    public Optional<Output<FileSourceRawArgs>> sourceRaw() {
        return Optional.ofNullable(this.sourceRaw);
    }

    /**
     * Timeout for uploading ISO/VSTMPL files in
     * seconds (defaults to 1800).
     * 
     */
    @Import(name="timeoutUpload")
    private @Nullable Output<Integer> timeoutUpload;

    /**
     * @return Timeout for uploading ISO/VSTMPL files in
     * seconds (defaults to 1800).
     * 
     */
    public Optional<Output<Integer>> timeoutUpload() {
        return Optional.ofNullable(this.timeoutUpload);
    }

    private FileArgs() {}

    private FileArgs(FileArgs $) {
        this.contentType = $.contentType;
        this.datastoreId = $.datastoreId;
        this.nodeName = $.nodeName;
        this.overwrite = $.overwrite;
        this.sourceFile = $.sourceFile;
        this.sourceRaw = $.sourceRaw;
        this.timeoutUpload = $.timeoutUpload;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FileArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FileArgs $;

        public Builder() {
            $ = new FileArgs();
        }

        public Builder(FileArgs defaults) {
            $ = new FileArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param contentType The content type. If not specified, the content
         * type will be inferred from the file extension. Valid values are:
         * 
         * @return builder
         * 
         */
        public Builder contentType(@Nullable Output<String> contentType) {
            $.contentType = contentType;
            return this;
        }

        /**
         * @param contentType The content type. If not specified, the content
         * type will be inferred from the file extension. Valid values are:
         * 
         * @return builder
         * 
         */
        public Builder contentType(String contentType) {
            return contentType(Output.of(contentType));
        }

        /**
         * @param datastoreId The datastore id.
         * 
         * @return builder
         * 
         */
        public Builder datastoreId(Output<String> datastoreId) {
            $.datastoreId = datastoreId;
            return this;
        }

        /**
         * @param datastoreId The datastore id.
         * 
         * @return builder
         * 
         */
        public Builder datastoreId(String datastoreId) {
            return datastoreId(Output.of(datastoreId));
        }

        /**
         * @param nodeName The node name.
         * 
         * @return builder
         * 
         */
        public Builder nodeName(Output<String> nodeName) {
            $.nodeName = nodeName;
            return this;
        }

        /**
         * @param nodeName The node name.
         * 
         * @return builder
         * 
         */
        public Builder nodeName(String nodeName) {
            return nodeName(Output.of(nodeName));
        }

        /**
         * @param overwrite Whether to overwrite an existing file (defaults to
         * `true`).
         * 
         * @return builder
         * 
         */
        public Builder overwrite(@Nullable Output<Boolean> overwrite) {
            $.overwrite = overwrite;
            return this;
        }

        /**
         * @param overwrite Whether to overwrite an existing file (defaults to
         * `true`).
         * 
         * @return builder
         * 
         */
        public Builder overwrite(Boolean overwrite) {
            return overwrite(Output.of(overwrite));
        }

        /**
         * @param sourceFile The source file (conflicts with `source_raw`),
         * could be a local file or a URL. If the source file is a URL, the file will
         * be downloaded and stored locally before uploading it to Proxmox VE.
         * 
         * @return builder
         * 
         */
        public Builder sourceFile(@Nullable Output<FileSourceFileArgs> sourceFile) {
            $.sourceFile = sourceFile;
            return this;
        }

        /**
         * @param sourceFile The source file (conflicts with `source_raw`),
         * could be a local file or a URL. If the source file is a URL, the file will
         * be downloaded and stored locally before uploading it to Proxmox VE.
         * 
         * @return builder
         * 
         */
        public Builder sourceFile(FileSourceFileArgs sourceFile) {
            return sourceFile(Output.of(sourceFile));
        }

        /**
         * @param sourceRaw The raw source (conflicts with `source_file`).
         * 
         * @return builder
         * 
         */
        public Builder sourceRaw(@Nullable Output<FileSourceRawArgs> sourceRaw) {
            $.sourceRaw = sourceRaw;
            return this;
        }

        /**
         * @param sourceRaw The raw source (conflicts with `source_file`).
         * 
         * @return builder
         * 
         */
        public Builder sourceRaw(FileSourceRawArgs sourceRaw) {
            return sourceRaw(Output.of(sourceRaw));
        }

        /**
         * @param timeoutUpload Timeout for uploading ISO/VSTMPL files in
         * seconds (defaults to 1800).
         * 
         * @return builder
         * 
         */
        public Builder timeoutUpload(@Nullable Output<Integer> timeoutUpload) {
            $.timeoutUpload = timeoutUpload;
            return this;
        }

        /**
         * @param timeoutUpload Timeout for uploading ISO/VSTMPL files in
         * seconds (defaults to 1800).
         * 
         * @return builder
         * 
         */
        public Builder timeoutUpload(Integer timeoutUpload) {
            return timeoutUpload(Output.of(timeoutUpload));
        }

        public FileArgs build() {
            $.datastoreId = Objects.requireNonNull($.datastoreId, "expected parameter 'datastoreId' to be non-null");
            $.nodeName = Objects.requireNonNull($.nodeName, "expected parameter 'nodeName' to be non-null");
            return $;
        }
    }

}
