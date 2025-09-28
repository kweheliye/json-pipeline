# JSON Pipeline

An **open-source Go tool** for processing massive structured datasets with streaming, filtering, and transformation, enabling efficient preparation of data for analytics and reporting.

Designed to handle massive datasets, from gigabytes to terabytes

The tool processes files in **5MB chunks**, streaming and converting data efficiently into **Parquet format**, enabling analytics workflows without loading the entire dataset into memory. 
On a 12-core machine, an 80GB dataset can be converted to Parquet in under 5 minutes.

The tool also supports **deeply nested JSON files**, allowing you to filter, flatten, and transform complex structures into analytics-ready formats.

---

## Key Features

### Implemented Features
- ✅ Stream large structured datasets efficiently for analytics
- ✅ Convert JSON/NDJSON input to **Parquet format**
- ✅ Handle gzip-compressed files automatically
- ✅ Chunked processing (default 5MB parts) for memory efficiency
- ✅ Filter records by field/value to extract relevant data
- ✅ Flatten and transform **deeply nested JSON structures**
- ✅ Parallel processing using worker pool for chunked data

### Planned Features
- ⏳ Support advanced filtering and expressions
- ⏳ Schema-based transformations
- ⏳ Additional output formats: CSV, Avro
- ⏳ Cloud input sources: S3, GCS, HTTP URLs
- ⏳ Parallel & distributed processing for high performance

---

## Quick Start

### Install
```bash
git clone https://github.com/yourusername/json-pipeline.git
cd json-pipeline
go build -o json-pipeline .
```

### Usage

#### Basic Conversion to Parquet
```bash
./json-pipeline --input data.json --output out.parquet
```

#### Filtered Conversion
```bash
./json-pipeline --input data.json.gz --field status --value active --output active.parquet
```

- Input can be `.json` or `.json.gz`
- Output is **Parquet**, optimized for analytics
- Supports **chunked processing (5MB default)**
- Supports **deeply nested JSON**, automatically flattened
- Apply **filters** to extract only relevant records

---

## Example

**Input (deep_nested.json):**
```json
{
  "id": 1,
  "status": "active",
  "details": {
    "score": 90,
    "preferences": {
      "color": "blue",
      "food": "pizza"
    }
  }
}
```

**Command:**
```bash
 ./json-pipeline run \                                                   
  --input data/example1.json \
  --output out.parquet \
  --cpuprofile cpu.prof \
  --memprofile mem.prof
```

**Output:**
- Parquet file with flattened fields: `id`, `status`, `details.score`, `details.preferences.color`, `details.preferences.food`


**Analyzing the profiles**

Requires Graphviz installed (go tool pprof -http=:8080 cpu.prof)

- go tool pprof cpu.prof
- (pprof) top
- (pprof) list main
- (pprof) web   
---

## Roadmap
- [ ] Configuration file support (YAML/TOML)
- [ ] Advanced filtering and expressions
- [ ] Schema-based transformations
- [ ] Additional output formats: CSV, Avro
- [ ] Cloud connectors (S3, GCS, HTTP URLs)
- [ ] Metrics, logging, and error handling improvements

---

## Contributing
Pull requests are welcome! Please open an issue first to discuss major changes.

---

## License
MIT License
