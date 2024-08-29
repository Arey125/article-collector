ALTER TABLE articles ADD COLUMN slug TEXT NOT NULL DEFAULT '';

UPDATE articles
    SET slug = replace(rtrim(link, '/'), rtrim(rtrim(link, '/'), replace(rtrim(link, '/'), '/', '')), '');
