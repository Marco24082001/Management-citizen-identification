<template>
    <div class="row">
        <div class="col-12">
            <card :title="table1.title" :subTitle="table1.subTitle">
                <el-table :data="filterTableData" style="width: 100%" max-height="60vh">
                  <el-table-column fixed label="Số" prop="id" width="120" sortable/>
                  <el-table-column label="Ảnh" width="150">
                    <template #default="scope">
                      <el-avatar :size="120" shape="square">
                        <img :src=scope.row.avatarUrl />
                      </el-avatar>
                    </template>
                  </el-table-column>
                  <el-table-column label="Vân tay" width="150">
                    <template #default="scope">
                      <el-avatar :size="120" shape="square">
                        <img :src=scope.row.signatureUrl />
                      </el-avatar>
                    </template>
                  </el-table-column>
                  <!-- <el-table-column label="AvatarUrl" prop="avatarUrl" width="150" sortable /> -->
                  <!-- <el-table-column label="SignatureUrl" prop="signatureUrl" width="150" sortable /> -->
                  <el-table-column label="Họ và tên" prop="name" width="150" sortable />
                  <el-table-column label="Ngày sinh" prop="dateOfbirth" width="150" sortable/>
                  <el-table-column label="Quốc tịch" prop="nationality" width="150" sortable/>
                  <el-table-column label="Quê quán" prop="placeOforigin" width="150" sortable/>
                  <el-table-column label="Nơi thường trú" prop="placeOfresidence" width="190" sortable/>
                  <el-table-column label="Đặc điểm nhận dạng" prop="personalIdentification" width="190" sortable/>
                  <el-table-column label="Người xác nhận" prop="personSignature" width="170" sortable/>
                  <el-table-column label="Ngày tạo" prop="createdDate" width="150" sortable/>
                  <el-table-column label="Có giá trị đến" prop="dateOfExpiry" width="150" sortable/>
                  <el-table-column fixed="right" width="215">
                      <template #header>
                      <el-input v-model="search" size="small" placeholder="Type to search" />
                      </template>
                      <template #default="scope">
                      <el-button size="small" @click="handleDialogEdit(scope.$index, scope.row)"
                          >Edit</el-button
                      >
                      <el-button
                          size="small"
                          type="danger"
                          @click="handleDelete(scope.$index, scope.row)"
                          >Delete</el-button
                      >
                      <el-button
                          size="small"
                          type="info"
                          @click="handleGetHistory(scope.$index, scope.row)"
                          >History</el-button
                      >
                      </template>
                  </el-table-column>
                </el-table>
                <el-button class="mt-4" style="width: 100%" @click="handleDialogCreate">Add Identications</el-button>
            </card>
        </div>
        <div class="col-12">
        <el-dialog v-model="dialogFormVisible" title="Identications Dialog">
          <div class="container-dialog">
            <el-form
            class="login-form"
            :model="model"
            :rules="rules"
            ref="form"
            @submit.enter.prevent="login"
            label-width="10rem"
            label-position="left"
            >
              <el-form-item prop="avatarUrl" label="Ảnh">
                <el-upload class="avatar-uploader" action="http://localhost:3000/identications/upload" :show-file-list="false"
                  :on-success="handleAvatarSuccess" :before-upload="beforePhotoUpload">
                  <img v-if="model.avatarUrl" :src="model.avatarUrl" class="avatar" />
                  <el-icon v-else class="avatar-uploader-icon">
                    <Plus />
                  </el-icon>
                </el-upload>
              </el-form-item>
              <el-form-item prop="signatureUrl" label="Vân tay">
                <el-upload class="avatar-uploader" action="http://localhost:3000/identications/upload" :show-file-list="false"
                  :on-success="handleSignatureSuccess" :before-upload="beforePhotoUpload">
                  <img v-if="model.signatureUrl" :src="model.signatureUrl" class="avatar" />
                  <el-icon v-else class="avatar-uploader-icon">
                    <Plus />
                  </el-icon>
                </el-upload>
              </el-form-item>
              <el-form-item prop="id" label="Số">
                <el-input v-model="model.id" placeholder="Nhập số Id">
                </el-input>
              </el-form-item>
              <el-form-item prop="name" label="Họ và tên">
                <el-input v-model="model.name" placeholder="Nhập họ và tên">
                </el-input>
              </el-form-item>
              <el-form-item prop="dateOfbirth" label="Ngày sinh">
                  <el-input v-model="model.dateOfbirth" placeholder="Input dateOfbirth">
                  </el-input>
              </el-form-item>
              <el-form-item prop="sex" label="Giới tính">
                <el-select v-model="model.sex" placeholder="Nhập giới tính">
                  <el-option 
                    v-for="item in sexes"
                    :key="item"
                    :label="item"
                    :value="item"
                  />
                </el-select>
              </el-form-item>
              <el-form-item prop="nationality" label="Quốc tịch">
                  <el-input v-model="model.nationality" placeholder="Nhập quốc tịch">
                  </el-input>
              </el-form-item>
              <el-form-item prop="placeOforigin" label="Quê quán">
                  <el-input v-model="model.placeOforigin" placeholder="Nhập quê quán">
                  </el-input>
              </el-form-item>
              <el-form-item prop="placeOfresidence" label="Nơi thường trú">
                  <el-input v-model="model.placeOfresidence" placeholder="Nhập nơi thường trú">
                  </el-input>
              </el-form-item>
              <el-form-item prop="dateOfExpiry" label="Có giá trị đến">
                  <el-input v-model="model.dateOfExpiry" placeholder="Có giá trị đến">
                  </el-input>
              </el-form-item>
              <el-form-item prop="personalIdentification" label="Đặc điểm nhận dạng">
                  <el-input v-model="model.personalIdentification" placeholder="Input personalIdentification">
                  </el-input>
              </el-form-item>
              <el-form-item prop="personSignature" label="Người xác nhận">
                  <el-input v-model="model.personSignature" placeholder="Input personSignature">
                  </el-input>
              </el-form-item>
            </el-form>
          </div>
          <template #footer>
            <span class="dialog-footer">
              <el-button
                v-if="dialogFormStatus === 'create'"
                @click="handleCreate"
                type="primary"
              >Create</el-button>
              <el-button
                v-else
                @click="handleEdit"
                type="primary"
              >Edit</el-button>
              <el-button @click="dialogFormVisible = false">Cancel</el-button>
            </span>
          </template>
        </el-dialog>
        <el-dialog v-model="dialogHistoryVisible" title="Identication History" width="90%">
          <el-table :data="historyIdentication" style="width: 100%" max-height="70vh">
            <el-table-column fixed label="No" prop="record.id" width="150" sortable/>
            <el-table-column label="Ảnh" prop="record.avatarUrl" width="150">
              <template #default="scope">
                <el-avatar :size="120" shape="square">
                  <img :src=scope.row.record.avatarUrl />
                </el-avatar>
              </template>
            </el-table-column>
            <el-table-column label="Vân tay" width="150">
              <template #default="scope">
                <el-avatar :size="120" shape="square">
                  <img :src=scope.row.record.signatureUrl />
                </el-avatar>
              </template>
            </el-table-column>
            <el-table-column label="Name" prop="record.name" width="150" sortable />
            <el-table-column label="Ngày sinh" prop="record.dateOfbirth" width="150" sortable/>
            <el-table-column label="Nationality" prop="record.nationality" width="150" sortable/>
            <el-table-column label="Place of origin" prop="record.placeOforigin" width="150" sortable/>
            <el-table-column label="Place of residence" prop="record.placeOfresidence" width="190" sortable/>
            <el-table-column label="Có giá trị đến" prop="record.dateOfExpiry" width="150" sortable/>
            <el-table-column label="Person Identification" prop="record.personalIdentification" width="190" sortable/>
            <el-table-column label="Có giá trị đến" prop="record.dateOfExpiry" width="150" sortable/>
            <el-table-column label="Người xác nhận" prop="record.personSignature" width="170" sortable/>
            <el-table-column label="Ngày tạo" prop="record.createdDate" width="150" sortable/>
            <el-table-column label="Ngày sử đổi" prop="timestamp" width="150" sortable/>
            <el-table-column label="Đã xóa" prop="isDelete" width="150" sortable/>
          </el-table>
        </el-dialog>
      </div>
    </div>
</template>

<script>
import IdenticationService from "@/services/identication/identication.services";
import { uploadImage } from "@/utils/cloudinaryUtils";
import { Plus } from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'



export default {
  components: {
    Plus,
  },
  data() {
    return {
      identicationsData: {
        rows: [],
      },
      historyIdentication: [],
      table1 : {
        title: "Identications list",
        subTitle: "All Identications",
      },
      search: "",
      dialogFormVisible: false,
      dialogHistoryVisible: false,
      dialogFormStatus: "create",
      rawFileImage: {
        avatarUrl: "",
        signatureUrl: "",
      },
      sexes: ["Nam", "Nữ"],
      model: {
        id: "",
        name: "",
        dateOfbirth: "",
        sex: "",
        nationality: "",
        placeOforigin: "",
        placeOfresidence: "",
        personalIdentification: "",
        personSignature: "",
        signatureUrl: "",
        avatarUrl: "",
        dateOfExpiry: "",
        createdDate: "",
        isDelete: "",
      },
      rules: {
        id: [
          { required: true, message: "Id is required", trigger: "blur" }
        ],
        name: [
          { required: true, message: "Name is required", trigger: "blur" }
        ],
        dateOfbirth: [
          { required: true, message: "DateOfbirth is required", trigger: "blur" }
        ],
        sex: [
          { required: true, message: "Sex is required", trigger: "blur" }
        ],
        nationality: [
          { required: true, message: "Nationality is required", trigger: "blur" }
        ],
        placeOforigin: [
          { required: true, message: "PlaceOforigin is required", trigger: "blur"}
        ],
        placeOforigin: [
          { required: true, message: "PlaceOforigin is required", trigger: "blur"}
        ],
        placeOfresidence: [
          { required: true, message: "PlaceOfresidence is required", trigger: "blur"}
        ],
        personalIdentification: [
          { required: true, message: "PlaceOforigin is required", trigger: "blur"}
        ],
        personSignature: [
          { required: true, message: "PersonSignature is required", trigger: "blur"}
        ],
        signatureUrl: [
          { required: true, message: "SignatureUrl is required", trigger: "blur"}
        ],
        avatarUrl: [
          { required: true, message: "AvatarUrl is required", trigger: "blur"}
        ],
        dateOfExpiry: [
          { required: true, message: "DateOfExpiry is required", trigger: "blur"}
        ]
      }
    }
  },
  async created() {
    const res = await IdenticationService.getAll();
    this.identicationsData.rows = res.data.response;
    console.log(this.identicationsData.rows);
    this.table1.subTitle = this.identicationsData.rows.length + " identications";
  },
  computed: {
    filterTableData() {
      return this.identicationsData.rows.filter(
        (data) =>
          !this.search ||
          data.id.toLowerCase().includes(this.search.toLowerCase())
      )
    }
  },
  methods: {
    resetModel: function() {
      this.model.name = "";
      this.model.maximum = 1;
    },
    handleDialogEdit: function(index, row) {
      this.dialogFormStatus = "edit";
      this.dialogFormVisible = true;
      this.editId = row.id;
      this.model = {...row};
    },
    handleDialogCreate: function() {
      this.dialogFormStatus = "create";
      this.dialogFormVisible = true;
      this.resetModel();
    },
    handleGetAll: async function() {
      const res = await IdenticationService.getAll();
      if (!res.data.error) {
        this.identicationsData.rows = res.data.response;
        console.log(this.identicationsData.rows);
      }
    },
    handleEdit: async function () {
      await this.$refs.form.validate(async (valid, fields) => {
        if (valid) {
          console.log("valid");
          console.log(this.model.avatarUrl);
          // try {
          //   const { url } = await uploadImage(this.rawFileImage.avatarUrl);
          //   this.model.avatarUrl = url;
          // } catch (error) {
          //   ElMessage.error('Upload image fail!');
          //   return
          // }
          const res = await IdenticationService.edit({ id: this.editId, ...this.model });
          console.log(res);
          if (!res.data.error) {
            this.handleGetAll();
            this.closedialogForm();
            ElMessage.success('Chỉnh sửa công cước công dân thành công!')
          }
        } else {
          console.log("invalid");
        }
      });
    },
    handleDelete: async function (index, row) {
      const res = await IdenticationService.delete(row.id);
      if (!res.data.error) {
        this.handleGetAll();
        ElMessage.success('Xóa công cước công dân thành công!')
      }
    },
    handleCreate: async function () {
      let valid = await this.$refs.form.validate();
      if (!valid) {
        console.log("invalid")
        return
      }
      const res = await IdenticationService.create(this.model);
      console.log(res.data.error)
      if (!res.data.error) {
        this.handleGetAll();
        this.closedialogForm();
        ElMessage.success('Tạo công cước công dân thành công!')
      }
    },
    handleGetHistory: async function(index, row) {
      const res = await IdenticationService.getHistory(row.id);
      if (!res.data.error) {
        this.historyIdentication = res.data.response;
      };
      console.log(res);
      this.dialogHistoryVisible = true;
    },
    getData: async function() {
      let responseData = await IdenticationService.getAll();
      return responseData;
    },
    closedialogForm: function() {
      this.$refs.form.resetFields();
      this.dialogFormVisible = false;
    },
    handleAvatarSuccess: async function (response, uploadFile) {
      this.model.avatarUrl = URL.createObjectURL(uploadFile.raw);
      this.rawFileImage.avatarUrl = uploadFile.raw;
      console.log(this.model);
    },
    handleSignatureSuccess: async function (response, uploadFile) {
      this.model.signatureUrl = URL.createObjectURL(uploadFile.raw);
      this.rawFileImage.signatureUrl = uploadFile.raw;
    },
    beforePhotoUpload: function (rawFile) {
      if (rawFile.type !== 'image/jpeg' && rawFile.type !== 'image/png') {
        ElMessage.error('Picture must be JPG|PNG format!');
        return false;
      } else if (rawFile.size / 1024 / 1024 > 2) {
        ElMessage.error('Picture size can not exceed 2MB!');
        return false;
      }
      return true;
    },
    errorHandler: () => true,
  },
}
</script>

<style scope>
.container-dialog {
  overflow-y: scroll;
  max-height:400px;
}

.avatar-uploader .avatar {
  width: 178px;
  height: 178px;
  display: block;
}
</style>

<style>
.avatar-uploader .el-upload {
  border: 1px dashed var(--el-border-color);
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
  transition: var(--el-transition-duration-fast);
}

.avatar-uploader .el-upload:hover {
  border-color: var(--el-color-primary);
}

.el-icon.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 178px;
  height: 178px;
  text-align: center;
}

</style>