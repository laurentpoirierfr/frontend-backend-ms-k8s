# README

## Database Schema

![Schema](./docs/assets/schema.png)


## Populate table menu_items

```bash
$ migrate create -ext sql -dir db/migrations -seq create_menu_items_populate
```