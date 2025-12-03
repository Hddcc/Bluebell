<template>
  <div class="content">
    <div class="left">
      <div class="container">
        <div class="post-header">
          <h4 class="con-title">{{post.title}}</h4>
          <button v-if="post.author_id == currentUserID" @click="deletePost" class="delete-btn">删除帖子</button>
        </div>
        <div class="l-container">
          <div class="con-info">{{post.content}}</div>
          <div class="user-btn">
            <span class="btn-item">
              <i class="iconfont icon-comment"></i>comment
            </span>
          </div>
        </div>
      </div>
    </div>
    <div class="right">
      <div class="topic-info">
        <h5 class="t-header"></h5>
        <div class="t-info">
          <a class="avatar"></a>
          <span class="topic-name">b/{{post.community_name}}</span>
        </div>
        <p class="t-desc">树洞 树洞 无限树洞的树洞</p>
        <ul class="t-num">
          <li class="t-num-item">
            <p class="number">5.2m</p>
            <span class="unit">Members</span>
          </li>
          <li class="t-num-item">
            <p class="number">5.2m</p>
            <span class="unit">Members</span>
          </li>
        </ul>
        <div class="date">Created Apr 10, 2008</div>
        <button class="topic-btn">JOIN</button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Content",
  data(){
    return {
      post:{},
      currentUserID: null, // Initialize currentUserID
    }
  },
  methods:{
    getPostDetail() {
      this.$axios({
        method: "get",
        url: "/post/"+ this.$route.params.id,
      })
        .then(response => {
          console.log(1, response.data);
          if (response.code == 1000) {
            this.post = response.data;
          } else {
            console.log(response.msg);
          }
        })
        .catch(error => {
          console.log(error);
        });
    },
    deletePost() {
        if (window.confirm("你确定要删除这篇帖子吗？此操作不可撤销！")) {
            this.$axios({
                method: "delete",
                url: "/post/" + this.post.id,
            })
            .then(response => {
                if (response.code == 1000) {
                    alert("帖子删除成功！");
                    this.$router.push("/"); // Navigate back to home page
                } else {
                    alert(response.msg || "删除失败！");
                }
            })
            .catch(error => {
                console.log(error);
                alert("删除过程中发生错误，请稍后再试。");
            });
        }
    },
    getCurrentUser() {
        const loginResult = JSON.parse(localStorage.getItem("loginResult"));
        if (loginResult && loginResult.user_id) {
            this.currentUserID = loginResult.user_id;
        }
    }
  },
  mounted: function() {
    this.getPostDetail();
    this.getCurrentUser();
  }
};
</script>

<style lang="less" scoped>
.content {
  max-width: 100%;
  box-sizing: border-box;
  display: flex;
  flex-direction: row;
  justify-content: center;
  margin: 0 auto;
  padding: 20px 24px;
  margin-top: 48px;
  .left {
    flex-grow: 1;
    max-width: 740px;
    border-radius: 4px;
    word-break: break-word;
    background: #ffffff;
    border: #edeff1;
    flex: 1;
    margin: 32px;
    margin-right: 12px;
    padding-bottom: 30px;
    position: relative;
    .container {
      width: 100%;
      height: auto;
      position: relative;
      padding: 20px; /* Added padding to container */
      .post-header {
        display: flex;
        justify-content: space-between;
        align-items: center;
        margin-bottom: 15px;
        .con-title {
          color: #000000;
          font-size: 24px; /* Made title more prominent */
          font-weight: bold;
          line-height: 28px;
          text-decoration: none;
          word-break: break-word;
        }
        .delete-btn {
          background-color: #dc3545; /* Red color for delete */
          color: white;
          border: none;
          padding: 8px 12px;
          border-radius: 4px;
          cursor: pointer;
          font-size: 14px;
          &:hover {
            background-color: #c82333;
          }
        }
      }
      .l-container {
        padding: 0; /* Removed unnecessary left padding */
        margin-left: 0; /* Removed unnecessary left margin */
        .con-info{
          margin: 25px 0;
          padding: 15px 0;
          border-bottom: 1px solid grey;
          font-size: 16px; /* Made content more readable */
          line-height: 1.5;
        }
        .user-btn {
          font-size: 12px;
          display: flex;
          display: -webkit-flex;
          .btn-item {
            display: flex;
            display: -webkit-flex;
            align-items: center;
            margin-right: 10px;
            .iconfont{
              margin-right: 4px;
            }
          }
        }
      }
    }
    .comment {
      width: 100%;
      height: auto;
      position: relative;
      .c-left {
        .line {
          border-right: 2px solid #edeff1;
          // width: 20px;
          height: 100%;
          position: absolute;
          left: 20px;
        }
        .c-arrow {
          display: flex;
          display: -webkit-flex;
          position: absolute;
          z-index: 2;
          flex-direction: column;
          left: 12px;
          background: #ffffff;
          padding-bottom: 5px;
        }
      }
      .c-right {
        margin-left: 40px;
        padding-right: 10px;
        .c-user-info {
          margin-bottom: 10px;
          .name {
            color: #1c1c1c;
            font-size: 12px;
            font-weight: 400;
            line-height: 16px;
          }
          .num {
            padding-left: 4px;
            font-size: 12px;
            font-weight: 400;
            line-height: 16px;
            color: #7c7c7c;
          }
        }
        .c-content {
          font-family: Noto Sans, Arial, sans-serif;
          font-size: 14px;
          font-weight: 400;
          line-height: 21px;
          word-break: break-word;
          color: rgb(26, 26, 27);
        }
      }
    }
  }
  .right {
    flex-grow: 0;
    width: 312px;
    margin-top: 32px;
    .topic-info {
      width: 100%;
      // padding: 12px;
      cursor: pointer;
      background-color: #ffffff;
      color: #1a1a1b;
      border: 1px solid #cccccc;
      border-radius: 4px;
      overflow: visible;
      word-wrap: break-word;
      padding-bottom: 30px;
      .t-header {
        width: 100%;
        height: 34px;
        background: #0079d3;
      }
      .t-info {
        padding: 0 12px;
        display: flex;
        display: -webkit-flex;
        width: 100%;
        height: 54px;
        align-items: center;
        .avatar {
          width: 54px;
          height: 54px;
          background: url("../assets/images/avatar.png") no-repeat;
          background-size: cover;
          margin-right: 10px;
        }
        .topic-name {
          height: 100%;
          line-height: 54px;
          font-size: 16px;
          font-weight: 500;
        }
      }
      .t-desc {
        font-family: Noto Sans, Arial, sans-serif;
        font-size: 14px;
        line-height: 21px;
        font-weight: 400;
        word-wrap: break-word;
        margin-bottom: 8px;
        padding: 0 12px;
      }
      .t-num {
        padding: 0 12px 20px 12px;
        display: flex;
        display: -webkit-flex;
        align-items: center;
        border-bottom: 1px solid #edeff1;
        .t-num-item {
          list-style: none;
          display: flex;
          display: -webkit-flex;
          flex-direction: column;
          width: 50%;
          .number {
            font-size: 16px;
            font-weight: 500;
            line-height: 20px;
          }
          .unit {
            font-size: 12px;
            font-weight: 500;
            line-height: 16px;
            word-break: break-word;
          }
        }
      }
      .date {
        font-family: Noto Sans, Arial, sans-serif;
        font-size: 14px;
        font-weight: 400;
        line-height: 18px;
        margin-top: 20px;
        padding: 0 12px;
      }
      .topic-btn {
        width: 286px;
        height: 34px;
        line-height: 34px;
        color: #ffffff;
        margin: 12px auto 0 auto;
        background: #003f6d;
        border-radius: 4px;
        box-sizing: border-box;
        margin-left: 13px;
      }
    }
  }
}
</style>