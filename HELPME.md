
## Example join with preload

```go
    query := p.db.Preload("Role").Model(&users.User{}).Where(users.User{RoleID: 2}, users.User{IsActive: true})
```


## Add pagination in responses

```go
        // default value limit
	limit := 25
	page := 1

	if len(queryPage) != 0 {
		i, _ := strconv.Atoi(queryPage)
		page = i
	}

	if len(queryLimit) != 0 {
		i, _ := strconv.Atoi(queryLimit)
		limit = i
	}

    .....

        var pages dto.Paginate
	offset := (page - 1) * limit
	tpages := float64(count) / float64(2)
	pages.TotalPage = math.Ceil(float64(tpages))
	pages.Page = 2
	pages.Count = count

	type responsewithpage struct {
		User     []dto.User   `json:"seller"`
		Paginate dto.Paginate `json:"pagination"`
	}

	res := responsewithpage{
		User:     user,
		Paginate: pages,
	}
	c.JSON(wrapper.StatusOK.New(constant.ApiSuccess, res))
```

## Create index using gist
```sql
CREATE INDEX  ON products USING gist (ll_to_earth(lat, long));
```



## Create trigger
```sql

CREATE OR REPLACE FUNCTION update_stocks() 
   RETURNS TRIGGER 
   LANGUAGE PLPGSQL
AS $$
BEGIN
   -- trigger logic
   UPDATE products SET stocks = stocks - NEW.quantity WHERE id = NEW.product_id AND
   stocks - NEW.quantity >= 1 ;
   RETURN NEW;
END;
$$;



CREATE TRIGGER
	stok_barang AFTER INSERT ON orders
	FOR EACH ROW EXECUTE PROCEDURE update_stocks();
```