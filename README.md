# NoiseAware Golang practice

## Testing
### Using `mockgen`

Template `mockgen -destination=<file to place mock into> -package=mock <Full package of interface to mock> <Interface name within package>`

Example `mockgen -destination=lib/message/repo/mock/mock_repo.go -package=mock github.com/tylerlofgren/noiseaware/lib/message/repo Repo`
