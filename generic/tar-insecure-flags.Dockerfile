# ruleid: tar-insecure-flags
RUN tar -xvf --absolute-paths archive.tar

# ruleid: tar-insecure-flags
RUN tar -xvf -P archive.tar

# ok: tar-insecure-flags
RUN tar -xvf --Psomeotherflag archive.tar

# ok: tar-insecure-flags
RUN tar -xvf archive.tar

# ok: tar-insecure-flags
RUN tar -xvf some_archive.tar
RUN wget "https://www.example.com/some/path" -P /tmp/some/local/path

# ok: tar-insecure-flags
RUN tar -xvf some_archive.tar && wget "https://www.example.com/some/path" -P /tmp/some/local/path