# go-prac

### Learnings
-   http.HandlerFunc for Handler Function on a specific route
-   http.NewServeMux() creates router instance and we can attach handler functions on it.
-   Handler Functions has access to w *http.ResponseWriter and r *http.Request, both respectivly are http Response Writer Pointer and Request Pointer just like res and req objects in express.js
-   