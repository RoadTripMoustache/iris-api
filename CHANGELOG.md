# CHANGELOG

## Unreleased
### Changed
- chore: Update dependencies

## 1.0.0 [02/11/2025]
### Added
- feat: GET /v1/ideas/{id} returns full idea details (including voters, comments and images)
- feat: Add endpoint to upload and get a image
- feat: Add endpoint to identify is the user is an admin
- feat: Add cron to remove orphan images
- ci: Create ci pipelines

### Changed
- feat: GET /v1/ideas no longer returns voters, comments or images fields in list items; response remains paginated with links+data