# Aspose.PDF Cloud SDK for Go — Agent Analysis

> **Repository:** [aspose-pdf-cloud/aspose-pdf-cloud-go](https://github.com/aspose-pdf-cloud/aspose-pdf-cloud-go)  
> **Version:** 26.7 | **Go Module:** `github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26`  
> **License:** MIT | **Go Version:** 1.16+  
> **API Version:** v3.0

---

## 1. Repository Overview

The **Aspose.PDF Cloud SDK for Go** is a generated REST API client that wraps the Aspose.PDF Cloud API v3.0. It enables Go applications to perform a wide range of PDF document processing operations — creation, manipulation, conversion, and rendering — entirely in the cloud.

The SDK is auto-generated from the OpenAPI specification and follows a **flat, single-package** structure (`package asposepdfcloud`). All types, services, and models reside in the root directory with no sub-packages.

---

## 2. Architecture & Core Components

### 2.1 Package Structure

```
asposepdfcloud/
├── api_client.go          # Core HTTP client, OAuth2 auth, request building
├── configuration.go       # Configuration struct & factory
├── pdf_api.go             # PdfApiService — 200+ API methods (~1.3 MB)
├── aspose_response.go     # Base response type
├── api_response.go        # HTTP response wrapper
├── base_test.go           # Test infrastructure singleton
├── {model}.go             # flat model files
├── {feature}_test.go      # test files
├── .devcontainer/         # Dev container configuration
├── docs/                  # Markdown API docs
├── settings/              # Credentials JSON
├── test_data/             # Test fixture PDF files
└── uses_cases/            # domain-specific runnable examples
```

### 2.2 Core Files

| File | Purpose |
|------|---------|
| **`api_client.go`** | HTTP client with OAuth2 client credentials flow, multipart upload, request preparation, response deserialization, cache control parsing |
| **`configuration.go`** | `Configuration` struct holding `BasePath`, `ClientId`, `ClientSecret`, `AccessToken`, `SelfHost` flag, custom `HTTPClient`, and default headers |
| **`pdf_api.go`** | `PdfApiService` — the main API surface with all REST endpoint methods (200+ methods, ~1.3 MB) |
| **`aspose_response.go`** | `AsposeResponse` — base response with `Code` (int32) and `Status` (string) |
| **`api_response.go`** | `APIResponse` — wraps `*http.Response` with `Message`, `Operation`, `RequestURL`, `Method`, `Payload` |

---

## 3. Data Model Organization

### 3.1 Model Files

All models are flat files in the root directory. Each PDF concept gets its own file:

| Category | Example Files |
|----------|---------------|
| **Annotations** | `annotation.go`, `annotation_type.go`, `annotation_flags.go`, `annotation_state.go`, `annotation_info.go`, `caret_annotation.go`, `circle_annotation.go`, `file_attachment_annotation.go`, `free_text_annotation.go`, `highlight_annotation.go`, `ink_annotation.go`, `line_annotation.go`, `link_annotation.go`, `movie_annotation.go`, `polygon_annotation.go`, `poly_line_annotation.go`, `popup_annotation.go`, `redaction_annotation.go`, `screen_annotation.go`, `sound_annotation.go`, `square_annotation.go`, `squiggly_annotation.go`, `stamp_annotation.go`, `strike_out_annotation.go`, `text_annotation.go`, `underline_annotation.go` |
| **Form Fields** | `field.go`, `field_type.go`, `form_field.go`, `check_box_field.go`, `combo_box_field.go`, `list_box_field.go`, `radio_button_field.go`, `text_box_field.go`, `signature_field.go`, `choice_field.go` |
| **Document** | `document.go`, `document_config.go`, `document_property.go`, `document_properties.go`, `display_properties.go`, `document_privilege.go` |
| **Pages** | `page.go`, `pages.go`, `page_layout.go`, `page_mode.go`, `page_range.go`, `page_word_count.go` |
| **Stamps** | `stamp.go`, `image_stamp.go`, `text_stamp.go`, `page_number_stamp.go`, `pdf_page_stamp.go`, `image_stamp_page_specified.go` |
| **Conversions** | `doc_format.go`, `html_document_type.go`, `epub_recognition_mode.go`, `color_depth.go`, `compression_type.go` |
| **Storage** | `file_version.go`, `file_versions.go`, `files_list.go`, `files_upload_result.go`, `disc_usage.go`, `object_exist.go` |
| **Primitives** | `color.go`, `point.go`, `rectangle.go`, `dash.go`, `border.go`, `border_info.go`, `margin_info.go`, `graph_info.go`, `link.go`, `link_element.go` |
| **Enums** | `annotation_type.go`, `border_style.go`, `border_effect.go`, `cap_style.go`, `direction.go`, `font_styles.go`, `horizontal_alignment.go`, `justification.go`, `line_ending.go`, `line_spacing.go`, `crypto_algorithm.go`, `permissions_flags.go` |

### 3.2 Response Type Naming Convention

- **Single entity:** `{Entity}Response` — e.g., `DocumentResponse`, `BookmarkResponse`, `CircleAnnotationResponse`
- **Collection:** `{Entity}sResponse` — e.g., `BookmarksResponse`, `CircleAnnotationsResponse`, `FieldsResponse`
- **Base:** `AsposeResponse` with `Code` (int32) and `Status` (string)

---

## 4. API Capabilities

### 4.1 Document Operations

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GetDocument` | GET `/pdf/{name}` | Read document info |
| `PutCreateDocument` | PUT `/pdf/{name}` | Create empty document |
| `PostCreateDocument` | POST `/pdf/{name}` | Create document with config |
| `PostOptimizeDocument` | POST `/pdf/{name}/optimize` | Optimize document (compress images, remove unused objects, unembed fonts) |
| `PostSplitDocument` | POST `/pdf/{name}/split` | Split document into pages |
| `PostSplitRangePdfDocument` | POST `/pdf/{name}/split/range` | Split by page ranges |
| `PostOrganizeDocument` | POST `/pdf/{name}/organize` | Reorder pages |
| `PostOrganizeDocuments` | POST `/pdf/organize` | Organize pages from multiple documents |
| `PostMergeDocuments` | PUT `/pdf/{name}/merge` | Merge multiple documents |

### 4.2 Page Operations

| Method | Description |
|--------|-------------|
| `GetPage` | Read page info |
| `PostPage` | Add new page |
| `DeletePage` | Delete page by number |
| `PostMovePage` | Move page to new position |
| `PostDocumentPagesRotate` | Rotate pages by angle |
| `PostDocumentPagesResize` | Resize pages |
| `PostDocumentPagesCrop` | Crop pages |
| `GetPageConvertToTiff/Jpeg/Png/Emf/Bmp/Gif` | Convert page to image (GET) |
| `PutPageConvertToTiff/Jpeg/Png/Emf/Bmp/Gif` | Convert page to image (PUT) |
| `PostPageImageStamps` | Add image stamp to page |
| `PostPageTextStamps` | Add text stamp to page |
| `PostPagePdfPageStamps` | Add PDF page stamp to page |
| `PostPagePageNumberStamps` | Add page number stamp |

### 4.3 Annotations (15+ Types)

Each annotation type supports full CRUD operations:

| Operation | Pattern |
|-----------|---------|
| **Get all** | `GetDocument{Type}Annotations(name, args)` |
| **Get by page** | `GetPage{Type}Annotations(name, pageNumber, args)` |
| **Get by ID** | `Get{Type}Annotation(name, annotationId, args)` |
| **Create** | `PostPage{Type}Annotations(name, pageNumber, annotation, args)` |
| **Update** | `Put{Type}Annotation(name, annotationId, annotation, args)` |
| **Delete** | `DeleteAnnotation(name, annotationId, args)` |
| **Flatten** | `PutAnnotationsFlatten(name, args)` |

Supported annotation types: Caret, Circle, FileAttachment, FreeText, Highlight, Ink, Line, Link, Movie, Polygon, PolyLine, Popup, Redaction, Screen, Sound, Square, Squiggly, Stamp, StrikeOut, Text, Underline.

### 4.4 Form Fields (8 Types)

| Field Type | Operations |
|------------|------------|
| CheckBox, ComboBox, ListBox, RadioButton, TextBox, Signature | Get document fields, get page fields, get by name, create, update, delete |
| General | `GetFields`, `PutUpdateFields`, `PutFieldsFlatten`, `PostFlattenDocument` |
| Import/Export | XML, FDF, XFDF formats (GET and PUT for each) |

### 4.5 Bookmarks

| Method | Description |
|--------|-------------|
| `GetDocumentBookmarks` | Get bookmark tree |
| `GetBookmarks` | Get bookmarks at path |
| `GetBookmark` | Get single bookmark |
| `PostBookmark` | Add bookmark |
| `PutBookmark` | Update bookmark |
| `DeleteBookmark` | Delete bookmark |
| `DeleteDocumentBookmarks` | Delete all bookmarks |

### 4.6 Conversions

**PDF → Other formats:**
DOC, DOCX, EPUB, Excel (XLS/XLSX), HTML, MobiXML, PDF/A, PPTX, SVG, TEX, TIFF, XLS, XML, XPS, and more.

**Other formats → PDF:**
APS, BMP, EPUB, GIF, HTML, JPEG, Markdown, MHTML, PCL, PNG, PS, SVG, TeX, Web, XML, XPS, XSL FO, images.

**Pattern:** `Get{Format}InStorageToPdf` / `Put{Format}InStorageToPdf` for each source format.

### 4.7 Storage & File Management

| Method | Description |
|--------|-------------|
| `UploadFile` | Upload file to cloud storage |
| `DownloadFile` | Download file from cloud storage |
| `CopyFile` / `MoveFile` / `DeleteFile` | File operations |
| `CreateFolder` / `CopyFolder` / `MoveFolder` / `DeleteFolder` | Folder operations |
| `GetFilesList` | List files in folder |
| `GetDiscUsage` | Get storage usage |
| `ObjectExists` / `StorageExists` | Check existence |
| `GetFileVersions` | List file versions |

### 4.8 Other Features

| Feature | Key Methods |
|---------|-------------|
| **Text** | `GetText`, `GetPageText`, `PutAddText` |
| **Images** | `GetImages`, `GetImage`, `DeleteImage`, `PostInsertImage` |
| **Links** | `GetPageLinkAnnotations`, `PostPageLinkAnnotations`, `PutLinkAnnotation`, `DeleteLinkAnnotation` |
| **Stamps** | `GetDocumentStamps`, `PostPageTextStamps`, `PostPageImageStamps`, `DeleteStamp` |
| **Tables** | `GetDocumentTables`, `PostPageTables`, `PutTable`, `DeleteTable` |
| **Watermarks** | Via image stamps |
| **Headers/Footers** | Via text/image stamps |
| **Encryption** | `PutEncryptDocument`, `PutDecryptDocument`, `PutChangePasswordDocument` |
| **Properties** | `GetDocumentProperties`, `PutSetProperty`, `DeleteProperty` |
| **XMP Metadata** | `GetXmpMetadataJson`, `GetXmpMetadataXml`, `PostXmpMetadata` |
| **Layers** | `GetDocumentLayers`, `DeleteDocumentLayer` |
| **Compare** | `PostCompareDocument` |
| **Privileges** | `PutPrivileges` |

---

## 5. Testing Infrastructure

### 5.1 Test Base (`base_test.go`)

- Singleton pattern via `GetBaseTest()` / `NewBaseTest()`
- Reads credentials from `settings/credentials.json`
- Supports both public cloud and self-hosted modes
- Provides `UploadFile(name)` helper
- Tracks test number via `GetTestNumber()`

### 5.2 Credentials Format

```json
{
    "api_url": "https://api.aspose.cloud/v3.0",
    "client_id": "YOUR_CLIENT_ID",
    "client_secret": "YOUR_CLIENT_SECRET",
    "self_host": false
}
```

### 5.3 Test Pattern

All tests follow a consistent pattern:

```go
func TestGetDocument(t *testing.T) {
    name := "4pages.pdf"
    if err := GetBaseTest().UploadFile(name); err != nil {
        t.Error(err)
    }
    args := map[string]interface{}{
        "folder": GetBaseTest().remoteFolder,
    }
    response, httpResponse, err := GetBaseTest().PdfAPI.GetDocument(name, args)
    if err != nil {
        t.Error(err)
    } else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
        t.Fail()
    } else {
        fmt.Printf("%d\tTestGetDocument - %d\n", GetBaseTest().GetTestNumber(), response.Code)
    }
}
```

## 6. Use Cases (`uses_cases/`)

The `uses_cases/` directory contains **domain-specific, runnable examples**:

### 6.1 Directory structure

Each domain follows a consistent structure:

```
uses_cases/{domain}/
├── {helper}_helper.go          # (optional) Shared utilities: initPdfApi(), uploadFile(), downloadFile()
├── {init}_init.go              # (optional) Data initialization: Not invoking of cloud methods
├── {main}_launch.go            # File(s) containing main() entry point
├── {operation_name}.go         # (optional) Individual domain operation that invokes cloud methods
└── {another_operation}.go      # (optional) Additional operations that invokes cloud methods
```

### 6.2 README.md Format

The `uses_cases/README.md` file serves as the **index and documentation** for all use case domains. It follows a strict format contains only sections:

#### Section Structure

Each domain is documented as a Markdown section with the following structure:

```markdown
#### {domain_name}
- **[{domain_name}/{main}.go]({domain_name}/{main}.go)** – Description of file contains `func main()` method.
  ```bash
  go run uses_cases/{domain_name}/*
  ```
- *[{domain_name}/operation1.go]({domain_name}/operation1.go)* – Description of the operation containing cloud method invocation.
- *[{domain_name}/operation2.go]({domain_name}/operation2.go)* – Description of the operation containing cloud method invocation.
```

**Note:** Domains with multiple main files list each one as a separate bold entry, each with its own `go run` command block.


#### Formatting Rules

| Element | Rule |
|---------|------|
| **Section header** | `#### {domain_name}` — level-4 heading, lowercase with underscores |
| **Main files** | Listed first, bold (`**`), includes `go run uses_cases/{domain_name}/*` command in a code block |
| **Operation files** | Listed after main files, italic (`*`), one per line |
| **File links** | Relative paths from `uses_cases/` directory, e.g., `[bookmarks/bookmarks_launch.go](bookmarks/bookmarks_launch.go)` (use actual file name) |
| **Descriptions** | Present tense, action-oriented (e.g., "Adds", "Retrieves", "Deletes") |
| **Blank lines** | One blank line between sections |

#### Example

```markdown
#### bookmarks
- **[bookmarks/bookmarks_launch.go](bookmarks/bookmarks_launch.go)** – Orchestrates bookmark CRUD operations including extraction.
  ```bash
  go run uses_cases/bookmarks/*
  ```
- *[bookmarks/append_bookmark.go](bookmarks/append_bookmark.go)* – Adds a new bookmark with specified properties (title, color, page, action) to a PDF document.
- *[bookmarks/get_bookmark_by_path_show.go](bookmarks/get_bookmark_by_path_show.go)* – Retrieves and displays bookmarks at a specific path in the bookmark hierarchy.
- *[bookmarks/get_bookmarks_show.go](bookmarks/get_bookmarks_show.go)* – Fetches and displays the entire bookmark tree of a PDF document.
- *[bookmarks/remove_bookmark.go](bookmarks/remove_bookmark.go)* – Deletes a bookmark at a specified path from a PDF document.
- *[bookmarks/replace_bookmark.go](bookmarks/replace_bookmark.go)* – Updates the properties of an existing bookmark at a specified path.
```

### 6.3 File Inclusion/Exclusion Rules

When generating or updating `uses_cases/README.md`, the following rules determine which files are included:

#### Included Files

| File Pattern | Reason | Example |
|-------------|--------|---------|
| `*{main}.go` (main files) | Files contains `func main()` function | `bookmarks_launch.go`, `stamps_launcher.go` |
| `*{operation}.go` (operation files) | Individual domain operations contains business logic (no `func main()` fucntion) | `append_bookmark.go`, `remove_bookmark.go` |

#### Excluded Files

| File Pattern | Reason | Example |
|-------------|--------|---------|
| **Non-`.go` files** | Only Go source files are documented | `*.md`, `*.json`, `*.pdf` |
| `*.go` (helper files) | Shared utilities (API init, file upload/download) | `bookmarks_helper.go` |
| `*.go` (data files)| Test data files (Initialization of complex data) | `fields_init.go` |
| **Files outside `uses_cases/`** | README only covers the `uses_cases/` directory | Root-level `*.go` files |
| **Hidden files/directories** | Not user-facing | `.DS_Store`, `.gitkeep` |

#### Ordering Rules

1. **Domains** are listed in **alphabetical order** by directory name.
2. **Within a domain**, files are ordered as:
   - `{main}.go` (first, bold)
   - All remaining `*.go` files in **alphabetical order** (italic)

---

## 7. Design Patterns & Conventions

### 7.1 Code Generation

The SDK is **auto-generated** from the OpenAPI specification. Evidence:
- `.swagger-codegen-ignore` file present
- Consistent, repetitive method structure across all 200+ API methods
- Flat file-per-model organization
- Uniform error handling and parameter validation

### 7.2 Key Conventions

| Convention | Description |
|------------|-------------|
| **Single package** | All code in `package asposepdfcloud` |
| **Flat structure** | No sub-packages; all files in root |
| **MIT license header** | Every `.go` file starts with the same license block |
| **Three-tuple returns** | `(Model, *http.Response, error)` for all API methods |
| **Optional params as map** | `map[string]interface{}` for optional query parameters |
| **Custom headers** | `x-aspose-client: go sdk`, `x-aspose-client-version: 26.7.0` |
| **Self-host support** | `SelfHost` flag skips OAuth2 authentication |
| **Zero external deps** | Only Go standard library used |

### 7.3 Error Handling

```go
// API errors return the HTTP response body in the error message
if localVarHttpResponse.StatusCode >= 300 {
    bodyBytes, _ := io.ReadAll(localVarHttpResponse.Body)
    return localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
}
```

---

## 8. Dependencies & Build

### 8.1 Dependencies

- **Zero external dependencies** — only Go standard library
- Go 1.16+ required

### 8.2 Installation

```bash
go get -u github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v26
```

## 9. Documentation

The `docs/` directory contains **Markdown files** with list of available apis and models:

- `PdfApi.md` — Full API method reference
- `{Model}.md` — One file per model type (e.g., `Document.md`, `Annotation.md`, `Bookmark.md`)
- `{Response}.md` — One file per response type (e.g., `DocumentResponse.md`, `AnnotationsResponse.md`)

---
