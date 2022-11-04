ALTER TABLE categories
ADD CONSTRAINT abbrname_transactiontype_unique UNIQUE (name_abbr, transaction_type);