# Test

## Integration test

- CI(Github action)で実行

### Command(手動)

#### 起動

```bash
docker-compose -f docker-compose.test.yml up
```

- 解説
  - [docker-compose.test.yml](../../docker-compose.test.yml)を使用して起動。
  - ***command***にて起動時に結合テスト実行。

#### 手動実行

```bash
docker-compose exec backend-test go test -v ./test/integration/... -tags=integration
```

#### Clean up

```bash
docker-compose -f docker-compose.test.yml down -v --remove-orphans
```
