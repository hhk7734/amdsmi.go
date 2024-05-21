```shell
diff --git a/pkg/rocm_smi/rocm_smi/rocm_smi.h b/pkg/rocm_smi/rocm_smi/rocm_smi.h
index e8db583..4d4d450 100644
--- a/pkg/rocm_smi/rocm_smi/rocm_smi.h
+++ b/pkg/rocm_smi/rocm_smi/rocm_smi.h
@@ -1143,7 +1143,9 @@ typedef struct {
 /**
  * @brief Opaque handle to function-support object
  */
-typedef struct rsmi_func_id_iter_handle * rsmi_func_id_iter_handle_t;
+typedef struct {
+    struct rsmi_func_id_iter_handle* handle;
+} rsmi_func_id_iter_handle_t;
```