# Configs

This documentation details all the configurations available.

## Firebase
### YAML Structure
```yaml
firebase:
    credentials_file_path: 
    project_id:
    mock:
        enabled:
        data_file_path:  
```

### Parameters
- `firebase.credentials_file_path` : *(Env variable : `FIREBASE_CREDENTIALS_FILE_PATH`)* - Path to the firebase credentials file
- `firebase.project_id` : *(Env variable : `FIREBASE_PROJECT_ID`)* - Project id of the firebase
- `firebase.mock.enabled` : *(Env variable : `FIREBASE_MOCK_ENABLED`)* - Set to `true` if you want to enable the mocks and do some specific tests.
- `firebase.mock.data_file_path` : *(Env variable : `FIREBASE_MOCK_DATA_PATH`)* - Path to the mock resources

---

## Database
```yaml
database:
  mock:
    enabled:
    data_folder_path:
  mongo:
    uri:
    name: 
```

### Parameters
- `database.mock.enabled` : *(Env variable : `DATABASE_MOCK_ENABLED`)* - Set to `true` if you want to enable the mocks and do some specific tests.
- `database.mock.data_file_path` : *(Env variable : `DATABASE_MOCK_DATA_PATH`)* - Path to the mock resources
- `database.mongo.uri` : *(Env variable : `DATABASE_MONGO_URI`)* - URI path of the database
- `database.mongo.name` : *(Env variable : `DATABASE_MONGO_NAME`)* - Name of the collection to use in Mongo.

---

## Server
```yaml
server:
  expose:
  metrics_expose:
  origins_allowed:
  images_dir:
  images_base_url:
```

### Parameters
- `server.expose` : *(Env variable : `SERVER_EXPOSE`)* - Port to expose the API. *Default 8080*
- `server.metrics_expose` : *(Env variable : `SERVER_METRICS_EXPOSE`)* - Port to expose the API metrics. *Default 2121*
- `server.origins_allowed` : List of the origins allowed to call the API
- `server.images_dir` : *(Env variable : `SERVER_IMAGES_DIR`)* - Folder where the images will be stored
- `server.images_base_url` : *(Env variable : `SERVER_IMAGES_BASE_URL`)* - Base path to call to get the images. By default need to be `/v1/images` if you use the API endpoint

---

## Images
```yaml
images:
  maxImagesPerIdea:
  maxImagesPerComment:
  maxSize:
  acceptedExtensions:
```

### Parameters
- `images.maxImagesPerIdea` : *(Env variable : `IMAGE_MAX_IMAGES_PER_IDEA`)* - Max number of images per idea
- `images.maxImagesPerComment` : *(Env variable : `IMAGE_MAX_IMAGES_PER_COMMENT`)* - Max number of images per comment
- `images.maxSize` : *(Env variable : `IMAGE_MAX_SIZE`)* - Image max site in bytes
- `images.acceptedExtensions` : List of file extensions allowed for file update.
