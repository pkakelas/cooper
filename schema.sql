CREATE TABLE crawled_document (
    id TEXT PRIMARY KEY NOT NULL,
    url TEXT NOT NULl,
    title TEXT NOT NULL,
    neighbors TEXT NOT NULL,
    termfreq TEXT NOT NULL
);

CREATE TABLE document_frequency (
    term INTEGER NOT NULL,
    occurencies TEXT NOT NULL,
    UNIQUE(term)
)