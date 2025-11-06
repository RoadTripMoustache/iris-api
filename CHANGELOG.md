# CHANGELOG

## Unreleased


## 1.1.0 [06/11/2025]
### Added
- docs: Added license
- feat: Add admin management endpoints
- feat: Add config to setup default admin accounts

### Changed
- chore: Update dependencies
- docs: Update documentation

## 1.0.0 [02/11/2025]
### Added
- feat: GET /v1/ideas/{id} returns full idea details (including voters, comments and images)
- feat: Add endpoint to upload and get a image
- feat: Add endpoint to identify is the user is an admin
- feat: Add cron to remove orphan images
- ci: Create ci pipelines

### Changed
- feat: GET /v1/ideas no longer returns voters, comments or images fields in list items; response remains paginated with links+data