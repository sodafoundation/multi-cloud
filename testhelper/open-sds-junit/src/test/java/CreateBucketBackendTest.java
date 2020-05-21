import com.google.gson.Gson;
import com.opensds.HttpHandler;
import com.opensds.jsonmodels.akskresponses.SignatureKey;
import com.opensds.jsonmodels.authtokensresponses.AuthTokenHolder;
import com.opensds.jsonmodels.inputs.addbackend.AddBackendInputHolder;
import com.opensds.jsonmodels.inputs.createbucket.CreateBucketFileInput;
import com.opensds.jsonmodels.tokensresponses.TokenHolder;
import com.opensds.jsonmodels.typesresponse.Type;
import com.opensds.jsonmodels.typesresponse.TypesHolder;
import com.opensds.utils.Constant;
import com.opensds.utils.Logger;
import com.opensds.utils.TextUtils;
import com.opensds.utils.Utils;
import okhttp3.Response;
import org.json.JSONArray;
import org.json.JSONException;
import org.json.JSONObject;
import org.json.XML;
import org.junit.jupiter.api.*;

import java.io.File;
import java.io.IOException;
import java.util.List;

import static org.junit.Assert.assertEquals;
import static org.junit.jupiter.api.Assertions.*;

// how to get POJO from any response JSON, use this site
// http://pojo.sodhanalibrary.com/

@TestMethodOrder(MethodOrderer.OrderAnnotation.class)
class CreateBucketBackendTest {

    public static AuthTokenHolder getAuthTokenHolder() {
        return mAuthTokenHolder;
    }

    public TypesHolder getTypesHolder() {
        return mTypesHolder;
    }

    public static HttpHandler getHttpHandler() {
        return mHttpHandler;
    }

    private static AuthTokenHolder mAuthTokenHolder = null;
    private static TypesHolder mTypesHolder = null;
    private static HttpHandler mHttpHandler = new HttpHandler();

    @org.junit.jupiter.api.BeforeAll
    static void setUp() {
        TokenHolder tokenHolder = getHttpHandler().loginAndGetToken();
        mAuthTokenHolder = getHttpHandler().getAuthToken(tokenHolder.getResponseHeaderSubjectToken());
        mTypesHolder = getHttpHandler().getTypes(getAuthTokenHolder().getResponseHeaderSubjectToken(),
                getAuthTokenHolder().getToken().getProject().getId());
    }

    @Test
    @Order(1)
    @DisplayName("Test creating bucket and backend on OPENSDS")
    public void testCreateBucketAndBackend() throws IOException, JSONException {
        // load input files for each type and create the backend
        for (Type t : getTypesHolder().getTypes()) {
            List<File> listOfIInputsForType =
                    Utils.listFilesMatchingBeginsWithPatternInPath(t.getName(),
                            Constant.CREATE_BUCKET_PATH);
            Gson gson = new Gson();
            // add the backend specified in each file
            for (File file : listOfIInputsForType) {
                String content = Utils.readFileContentsAsString(file);
                assertNotNull(content);

                AddBackendInputHolder inputHolder = gson.fromJson(content, AddBackendInputHolder.class);
                int code = getHttpHandler().addBackend(getAuthTokenHolder().getResponseHeaderSubjectToken(),
                        getAuthTokenHolder().getToken().getProject().getId(),
                        inputHolder);
                assertEquals(code, 200);

                // backend added, now create buckets
                List<File> listOfIBucketInputs =
                        Utils.listFilesMatchingBeginsWithPatternInPath("bucket",
                                Constant.CREATE_BUCKET_PATH);
                SignatureKey signatureKey = getHttpHandler().getAkSkList(getAuthTokenHolder().getResponseHeaderSubjectToken(),
                        getAuthTokenHolder().getToken().getProject().getId());
                // create the bucket specified in each file
                for (File bucketFile : listOfIBucketInputs) {
                    String bucketContent = Utils.readFileContentsAsString(bucketFile);
                    assertNotNull(bucketContent);

                    CreateBucketFileInput bfi = gson.fromJson(bucketContent, CreateBucketFileInput.class);

                    // filename format is "bucket_<bucketname>.json", get the bucket name here
                    String bName = Utils.getBucketName(bucketFile);

                    // now create buckets
                    int cbCode = getHttpHandler().createBucket(bfi, bName, signatureKey);
                    assertEquals(cbCode, 200);
                    boolean isBucketExist = testGetListBuckets(bName, signatureKey);
                    assertTrue(isBucketExist, "Bucket is not exist: ");
                }
            }
        }
    }

    /**
     * Get bucket list.
     *
     * @param bName Bucket name
     * @param signatureKey Signature key object
     * @return boolean If bucket is exist return true else false
     * @throws JSONException Json exception
     * @throws IOException Io exception
     */
    private boolean testGetListBuckets(String bName, SignatureKey signatureKey)
            throws JSONException, IOException {
        Response listBucketResponse = getHttpHandler().getBuckets(signatureKey);
        int resCode = listBucketResponse.code();
        String responseBody = listBucketResponse.body().string();
        Logger.logString("Response Code: " + resCode);
        Logger.logString("Response: " + responseBody);
        JSONObject jsonObject = XML.toJSONObject(responseBody);
        JSONArray jsonObjectBucketList = jsonObject.getJSONObject("ListAllMyBucketsResult")
                .getJSONObject("Buckets").getJSONArray("Bucket");
        boolean isBucketExist = false;
        for (int i = 0; i < jsonObjectBucketList.length(); i++) {
            String bucketName = jsonObjectBucketList.getJSONObject(i).get("Name").toString();
            if (!TextUtils.isEmpty(bucketName)) {
                if (bucketName.equals(bName)) {
                    isBucketExist = true;
                }
            }
        }
        return isBucketExist;
    }

    @Test
    @Order(2)
    @DisplayName("Test creating bucket using Invalid name")
    public void testCreateBucketUsingCapsName() {
        // load input files for each type and create the backend
        for (Type t : getTypesHolder().getTypes()) {
            List<File> listOfIInputsForType =
                    Utils.listFilesMatchingBeginsWithPatternInPath(t.getName(),
                            Constant.CREATE_BUCKET_PATH);
            Gson gson = new Gson();
            // add the backend specified in each file
            for (File file : listOfIInputsForType) {
                String content = Utils.readFileContentsAsString(file);
                assertNotNull(content);

                // backend added, now create buckets
                List<File> listOfIBucketInputs =
                        Utils.listFilesMatchingBeginsWithPatternInPath("bucket",
                                Constant.CREATE_BUCKET_PATH);
                SignatureKey signatureKey = getHttpHandler().getAkSkList(getAuthTokenHolder().getResponseHeaderSubjectToken(),
                        getAuthTokenHolder().getToken().getProject().getId());
                // create the bucket specified in each file
                for (File bucketFile : listOfIBucketInputs) {
                    String bucketContent = Utils.readFileContentsAsString(bucketFile);
                    assertNotNull(bucketContent);

                    // filename format is "bucket_<bucketname>.json", get the bucket name here
                    CreateBucketFileInput bfi = gson.fromJson(bucketContent, CreateBucketFileInput.class);

                    // now create buckets
                    int cbCode = getHttpHandler().createBucket(bfi, "RATR_@#", signatureKey);
                    assertEquals(cbCode, 400);
                }
            }
        }
    }

    @Test
    @Order(3)
    @DisplayName("Test re-creating backend with same name on OPENSDS")
    public void testReCreateBackend() {
        // load input files for each type and create the backend
        for (Type t : getTypesHolder().getTypes()) {
            List<File> listOfIInputsForType =
                    Utils.listFilesMatchingBeginsWithPatternInPath(t.getName(),
                            Constant.CREATE_BUCKET_PATH);
            Gson gson = new Gson();
            // Re-create backend specified in each file
            for (File file : listOfIInputsForType) {
                String content = Utils.readFileContentsAsString(file);
                assertNotNull(content);

                AddBackendInputHolder inputHolder = gson.fromJson(content, AddBackendInputHolder.class);
                int code = getHttpHandler().addBackend(getAuthTokenHolder().getResponseHeaderSubjectToken(),
                        getAuthTokenHolder().getToken().getProject().getId(),
                        inputHolder);
                assertEquals("Re-create backend with same name:Response code not matched:",code, 409);
            }
        }
    }

    @Test
    @Order(4)
    @DisplayName("Test re-creating bucket with same name on OPENSDS")
    public void testReCreateBucket() {
        // load input files for each type and create the backend
        for (Type t : getTypesHolder().getTypes()) {
            List<File> listOfIInputsForType =
                    Utils.listFilesMatchingBeginsWithPatternInPath(t.getName(),
                            Constant.CREATE_BUCKET_PATH);
            Gson gson = new Gson();
            // Re-create backend specified in each file
            for (File file : listOfIInputsForType) {
                String content = Utils.readFileContentsAsString(file);
                assertNotNull(content);
                List<File> listOfIBucketInputs =
                        Utils.listFilesMatchingBeginsWithPatternInPath("bucket",
                                Constant.CREATE_BUCKET_PATH);
                SignatureKey signatureKey = getHttpHandler().getAkSkList(getAuthTokenHolder().getResponseHeaderSubjectToken(),
                        getAuthTokenHolder().getToken().getProject().getId());
                // create the bucket specified in each file
                for (File bucketFile : listOfIBucketInputs) {
                    String bucketContent = Utils.readFileContentsAsString(bucketFile);
                    assertNotNull(bucketContent);

                    CreateBucketFileInput bfi = gson.fromJson(bucketContent, CreateBucketFileInput.class);

                    // filename format is "bucket_<bucketname>.json", get the bucket name here
                    String bName = bucketFile.getName().substring(bucketFile.getName().indexOf("_") + 1,
                            bucketFile.getName().indexOf("."));

                    // now create buckets
                    int cbCode = getHttpHandler().createBucket(bfi, bName, signatureKey);
                    assertEquals("Re-create bucket with same name failed:Response code not matched: "
                            , cbCode, 409);
                }
            }
        }
    }

    @Test
    @Order(5)
    @DisplayName("Test create bucket with empty name")
    public void testCreateBucketWithEmptyName() throws IOException, JSONException {
        System.out.println("Verifying response code: Input (Backend name) with empty value in payload and bucket" +
                " name is empty");
        SignatureKey signatureKey = getHttpHandler().getAkSkList(getAuthTokenHolder().getResponseHeaderSubjectToken(),
                getAuthTokenHolder().getToken().getProject().getId());
        File bucketFile = new File(Constant.EMPTY_FIELD_PATH, "bucket_emptyvalue.json");
        String bucketContent = Utils.readFileContentsAsString(bucketFile);
        assertNotNull(bucketContent);
        Gson gson = new Gson();
        CreateBucketFileInput bfi = gson.fromJson(bucketContent, CreateBucketFileInput.class);
        int cbCode = getHttpHandler().createBucket(bfi, "", signatureKey);
        assertEquals("Bucket name and backend name is empty in payload :Response code not matched: "
                , cbCode, 405);
        boolean isBucketExist = testGetListBuckets("", signatureKey);
        assertFalse(isBucketExist);

        File file = new File(Constant.EMPTY_FIELD_PATH, "bucket_b1324.json");
        String content = Utils.readFileContentsAsString(file);
        assertNotNull(content);
        String bName = file.getName().substring(bucketFile.getName().indexOf("_") + 1,
                bucketFile.getName().indexOf("."));

        CreateBucketFileInput input = gson.fromJson(content, CreateBucketFileInput.class);
        int code = getHttpHandler().createBucket(input, bName, signatureKey);
        System.out.println("Verifying response code: In input (Backend name) with not valid value but bucket name is valid");
        assertEquals("Backend does not exist:Response code not matched: "
                , code, 404);
    }

    @Test
    @Order(6)
    @DisplayName("Test request body with empty value,try to create backend")
    public void testRequestBodyWithEmptyFieldBackend() {
        Gson gson = new Gson();
        File file = new File(Constant.EMPTY_FIELD_PATH+"ibm-cos_b1321.json");
        String content = Utils.readFileContentsAsString(file);
        assertNotNull(content);
        AddBackendInputHolder inputHolder = gson.fromJson(content, AddBackendInputHolder.class);
        int code = getHttpHandler().addBackend(getAuthTokenHolder().getResponseHeaderSubjectToken(),
                getAuthTokenHolder().getToken().getProject().getId(),
                inputHolder);
        Logger.logObject("Backend Input: "+content);
        assertEquals("Request body with empty value:Response code not matched:",code, 400);
    }

    @Test
    @Order(7)
    @DisplayName("Test uploading object in a bucket")
    public void testUploadObject() throws IOException, JSONException {
        // load input files for each type
        for (Type t : getTypesHolder().getTypes()) {
            List<File> listOfIInputsForType =
                    Utils.listFilesMatchingBeginsWithPatternInPath(t.getName(),
                            Constant.CREATE_BUCKET_PATH);
            for (File file : listOfIInputsForType) {
                String content = Utils.readFileContentsAsString(file);
                assertNotNull(content);
                List<File> listOfIBucketInputs =
                        Utils.listFilesMatchingBeginsWithPatternInPath("bucket",
                                Constant.CREATE_BUCKET_PATH);
                // Get bucket name.
                for (File bucketFile : listOfIBucketInputs) {
                    String bucketContent = Utils.readFileContentsAsString(bucketFile);
                    assertNotNull(bucketContent);
                    String bucketName = bucketFile.getName().substring(bucketFile.getName().indexOf("_") + 1,
                            bucketFile.getName().indexOf("."));
                    // Get object for upload.
                    File fileRawData = new File(Constant.RAW_DATA_PATH);
                    File[] files = fileRawData.listFiles();
                    String mFileName = null;
                    File mFilePath = null;
                    for (File fileName : files) {
                        mFileName = fileName.getName();
                        mFilePath = fileName;

                        SignatureKey signatureKey = getHttpHandler().getAkSkList(getAuthTokenHolder().getResponseHeaderSubjectToken(),
                                getAuthTokenHolder().getToken().getProject().getId());
                        int cbCode = getHttpHandler().uploadObject(signatureKey,
                                bucketName, mFileName, mFilePath);
                        assertEquals("Uploaded object failed", cbCode, 200);

                        //Verifying object is uploaded in bucket.
                        boolean isUploaded = testGetListOfObjectFromBucket(bucketName, mFileName, signatureKey);
                        assertTrue(isUploaded,"Object is not uploaded");
                    }
                }
            }
        }
    }

    /**
     * Get List of object from bucket
     *
     * @param bucketName   bucket name
     * @param fileName     file name
     * @param signatureKey signature key object
     * @return boolean object is upload then true else false
     * @throws JSONException json exception
     * @throws IOException   io exception
     */
    private boolean testGetListOfObjectFromBucket(String bucketName, String fileName,
                                                  SignatureKey signatureKey)
            throws JSONException, IOException {
        Response listObjectResponse = getHttpHandler().getBucketObjects(bucketName, signatureKey);
        int resCode = listObjectResponse.code();
        String resBody = listObjectResponse.body().string();
        Logger.logString("Response Code: " + resCode);
        Logger.logString("Response: " + resBody);
        assertEquals("Get list of object failed", resCode, 200);
        JSONObject jsonObject = XML.toJSONObject(resBody);
        JSONObject jsonObjectListBucket = jsonObject.getJSONObject("ListBucketResult");
        boolean isUploaded = false;
        if (jsonObjectListBucket.has("Contents")) {
            if (jsonObjectListBucket.get("Contents") instanceof JSONArray) {
                JSONArray objects = jsonObjectListBucket.getJSONArray("Contents");
                for (int i = 0; i < objects.length(); i++) {
                    if (!TextUtils.isEmpty(objects.getJSONObject(i).get("Key").toString())) {
                        if (objects.getJSONObject(i).get("Key").toString().equals(fileName)) {
                            isUploaded = true;
                        }
                    }
                }
            } else {
                if (!TextUtils.isEmpty(jsonObjectListBucket.getJSONObject("Contents")
                        .get("Key").toString())) {
                    if (jsonObjectListBucket.getJSONObject("Contents").get("Key").toString()
                            .equals(fileName)) {
                        isUploaded = true;
                    }
                }
            }
        }
        return isUploaded;
    }

    @Test
    @Order(8)
    @DisplayName("Test uploading object failed scenario")
    public void testUploadObjectFailed(){
        File fileRawData = new File(Constant.RAW_DATA_PATH);
        File[] files = fileRawData.listFiles();
        String mFileName = null;
        File mFilePath = null;
        for (File fileName : files) {
            mFileName = fileName.getName();
            mFilePath = fileName;
        }
        SignatureKey signatureKey = getHttpHandler().getAkSkList(getAuthTokenHolder().getResponseHeaderSubjectToken(),
                getAuthTokenHolder().getToken().getProject().getId());
        int cbCode = getHttpHandler().uploadObject(signatureKey,
                "bucketName", mFileName, mFilePath);
        System.out.println("Verifying Upload object with non existing bucket");
        assertEquals("Upload object with non existing bucket: Response code not matched", cbCode, 404);

        List<File> listOfIBucketInputs =
                Utils.listFilesMatchingBeginsWithPatternInPath("bucket",
                        Constant.CREATE_BUCKET_PATH);
        // Get bucket name.
        for (File bucketFile : listOfIBucketInputs) {
            String bucketName = bucketFile.getName().substring(bucketFile.getName().indexOf("_") + 1,
                    bucketFile.getName().indexOf("."));
            int code = getHttpHandler().uploadObject(signatureKey,
                    bucketName, "", mFilePath);
            System.out.println("Verifying upload object in existing bucket with file name is empty");
            assertEquals("Upload object with existing bucket with file name empty: Response code not matched"
                    , code, 400);
        }
    }

    @Test
    @Order(9)
    @DisplayName("Test verifying download non exist file")
    public void testDownloadNonExistFile() throws IOException {
        List<File> listOfIBucketInputs =
                Utils.listFilesMatchingBeginsWithPatternInPath("bucket",
                        Constant.CREATE_BUCKET_PATH);
        SignatureKey signatureKey = getHttpHandler().getAkSkList(getAuthTokenHolder().getResponseHeaderSubjectToken(),
                getAuthTokenHolder().getToken().getProject().getId());
        for (File bucketFile : listOfIBucketInputs) {
            String bucketContent = Utils.readFileContentsAsString(bucketFile);
            assertNotNull(bucketContent);
            String fileName = "download_image.jpg";
            String bucketName = bucketFile.getName().substring(bucketFile.getName().indexOf("_") + 1,
                    bucketFile.getName().indexOf("."));
            Response response = getHttpHandler().downloadObject(signatureKey,
                    bucketName, "23455", fileName);
            int code = response.code();
            String body = response.body().string();
            Logger.logString("Response Code: " + code);
            Logger.logString("Response: " + body);
            assertEquals("Downloading non exist file: ", code, 404);
        }
    }

    @Test
    @Order(10)
    @DisplayName("Test verifying download file from non exist bucket")
    public void testDownloadFileFromNonExistBucket() throws IOException {
        String dFileName = "download_image.jpg";
        File fileRawData = new File(Constant.RAW_DATA_PATH);
        File[] files = fileRawData.listFiles();
        String mFileName = null;
        for (File fileName : files) {
            mFileName = fileName.getName();
        }
        SignatureKey signatureKey = getHttpHandler().getAkSkList(getAuthTokenHolder().getResponseHeaderSubjectToken(),
                getAuthTokenHolder().getToken().getProject().getId());
        Response response = getHttpHandler().downloadObject(signatureKey,
                "hfhfhd", mFileName, dFileName);
        int code = response.code();
        String body = response.body().string();
        Logger.logString("Response Code: " + code);
        Logger.logString("Response: " + body);
        assertEquals("Downloading file from non exist bucket: ", code, 404);
    }

    @Test
    @Order(11)
    @DisplayName("Test verifying download file from non exist bucket and file name")
    public void testDownloadNonExistBucketAndFile() throws IOException {
        String fileName = "download_image.jpg";
        System.out.println("Verifying download file from non exist bucket and file name");
        SignatureKey signatureKey = getHttpHandler().getAkSkList(getAuthTokenHolder().getResponseHeaderSubjectToken(),
                getAuthTokenHolder().getToken().getProject().getId());
        Response response = getHttpHandler().downloadObject(signatureKey,
                "ghjhb", "yuyiyh", fileName);
        int code = response.code();
        String body = response.body().string();
        Logger.logString("Response Code: " + code);
        Logger.logString("Response: " + body);
        assertEquals("Downloading file from non exist bucket and file name: ", code, 404);
        assertEquals("Response message is not valid, bucket and filename not exist: ", body);
    }
}

