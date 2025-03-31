<template>
    <div>
      <el-table style="width: 100%; max-width: 1200px;" :data="tableData">
          <el-table-column type="expand">
              <template slot-scope="props">
                  <el-form label-position="left" inline class="demo-table-expand">
                      <el-form-item label="影视封面:">
                          <span>{{ props.row.VedioCover }}</span>
                      </el-form-item>
                      <el-form-item label="影视简介:">
                          <div class="discrip">
                            {{  props.row.Description }}
                          </div>
                      </el-form-item>
                      <el-form-item label="每集信息">
                          <el-collapse style="width: 60vw;">
                              <el-collapse-item name="1">
                                  <div v-for="(pr, index) in props.row.EpisodeInfos" :key="pr + index + 'index'"
                                      style="margin-top: 4px;">
                                      <el-col :span="4">
                                          第 {{ pr.Num }} 集 :
                                      </el-col>
                                      <el-col :span="20">
                                          <el-link @click="jumpToVideo()" target="_blank">{{ pr.Name }}</el-link>
                                      </el-col>
                                  </div>
                              </el-collapse-item>
                          </el-collapse>
                      </el-form-item>
                  </el-form>
              </template>
          </el-table-column>
          <el-table-column label="影视 ID" prop="VedioId" width="80"></el-table-column>
          <el-table-column label="影视名称" prop="VedioName" width="300"></el-table-column>
          <el-table-column label="影视类别" prop="TypeId">
            <template slot-scope="scope">
                <el-tag>{{ TypeList[scope.row.TypeId - 1]["Name"] }}</el-tag>
            </template>
        </el-table-column>
          <el-table-column label="前台可见性" prop="Visable">
            <template slot-scope="scope">
                <el-tag v-if="scope.row.Visable">可见</el-tag>
                <el-tag v-else>不可见</el-tag>
            </template>
          </el-table-column>
          <el-table-column label="总共集数" prop="Episode" width="100"></el-table-column>
          <el-table-column align="right">
              <template slot="header" slot-scope="scope">
                  <el-input v-model="search" size="mini" placeholder="输入关键字搜索"/>
              </template>
              <template slot-scope="scope">
                  <el-button
                      size="mini"
                      @click="handleEdit(scope.$index, scope.row)">Edit</el-button>
                  <el-button
                      size="mini"
                      type="danger"
                      @click="handleDelete(scope.$index, scope.row)">Delete</el-button>
              </template>     
          </el-table-column>
      </el-table>
      <Vedio_editDrawer ref="drawer"></Vedio_editDrawer>
    </div>
  </template>
  
  <script>
  import Vedio_editDrawer from './Vedio_editDrawer.vue';
  export default {
      components : {
        Vedio_editDrawer
      },
      data() {
          return {
              search : "",
              tableData: [],
              TypeList : []
          }
      },
      methods : {
          handleEdit(index, row) {
              this.$refs.drawer.handleOpen(JSON.parse(JSON.stringify(row)))
          },
          handleDelete(index, row) {
              this.$confirm('此操作将永久删除该影视（所有视频）, 是否继续?', '提示', {
                  confirmButtonText: '确定',
                  cancelButtonText: '取消',
                  type: 'warning'
              }).then(() => {
                // 这里插入删除的代码
                  this.$message({ type: 'success', message: '删除成功!'});
              }).catch(() => {
                  this.$message({ type: 'info', message: '已取消删除' });          
              });
          },
          jumpToVideo() {

          }
      },
      beforeCreate() {
        this.$api.admin.VedioType(this).then(res => {
          this.TypeList = res
        })        
        this.$api.admin.showVedio(this).then(res => {
            this.tableData = res
        })
      }
  }
  </script>
  
  <style scope>
  .demo-table-expand {
      font-size: 0;
      width: 100%;
  }
  .demo-table-expand label {
      width: 90px;
      color: #99a9bf;
      margin-left: 16px;
  }
  .demo-table-expand .el-form-item {
      margin-right: 0;
      margin-bottom: 0;
      width: 100%;
  }
  .discrip {
      width: 70% !important;
      margin-bottom : 8px;
      margin-left: 36px;
      line-height: 24px;
  }
  </style>
  