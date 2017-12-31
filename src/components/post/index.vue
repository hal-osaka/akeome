<template>
  <section>

    <div class="yakatabune">
      <img src="https://s3-us-west-2.amazonaws.com/akeome/yakatabune.png" alt="yakatabune">
    </div>

    <h1>お正月メッセージを送ろう！</h1>

    <div class="container">

      <div class="icon">
        <img :src="icon" alt="icon">
      </div>

      <p class="name">To {{nickname}}</p>

      <b-field horizontal>
        <textarea class="textarea" v-model="input"></textarea>
      </b-field>

      <b-field horizontal>
        <p class="control">
          <button class="button is-primary" @click="send">送る</button>
        </p>
      </b-field>

      <ul class="list">
        <li v-for="message in messages">
          <div v-if="message.body.length > 0">
            <p class="text">{{message.body}}</p>
            <p class="time">{{message.created_at}}</p>
          </div>
        </li>
      </ul>

      <div class="images">
        <div class="img">
          <img src="https://s3-us-west-2.amazonaws.com/akeome/hime.png" alt="hime">
        </div>
        <div class="img">
          <img src="https://s3-us-west-2.amazonaws.com/akeome/tono.png" alt="tono">
        </div>
      </div>
    </div>

    <b-loading :active.sync="isLoading" :canCancel="true"></b-loading>
  </section>
</template>

<script>
import HttpService from '../../service/http'

export default {
  name: "post",
  data() {
    return {
      icon: "https://screenshotlayer.com/images/assets/placeholder.png",
      nickname: "",
      messages: "",
      input: "",
      isLoading: true
    }
  },
  mounted() {
    HttpService.GetProfile()
    .then((res) => {
      const { icon, nickname } = res.data
      if(nickname === "") {
        location.href = "/"
      }
      this.icon = icon
      this.nickname = nickname
    })
    .catch((err) => {
      console.error(err)
    })

    this.getMessages()
  },
  methods: {
    send() {
      const text = this.input
      if(text === "") {
        this.$toast.open({
          message: "メッセージが未入力です！",
          type: "is-danger"
        })
        return
      }
      HttpService.PostMessage(text)
      .then((res) => {
        this.$toast.open({
          message: "メッセージしました！",
          type: "is-success"
        })
        this.getMessages()
      })
      .catch((err) => {
        console.error(err)
      })
      this.input = ""
    },
    getMessages() {
      HttpService.GetMessages()
      .then((res) => {
        this.messages = res.data.messages
        this.isLoading = false
      })
      .catch((err) => {
        console.error(err)
      })
    }
  }
}
</script>

<style lang="scss" scoped>
h1 {
  font-size: 1.1rem;
  font-weight: bold;
  text-align: center;
  margin-top: 10px;
  margin-bottom: 30px;
}

.icon {
  width: 40vw;
  height: 40vw;
  margin: 0 auto;
  border-radius: 50%;
  overflow: hidden;
  img {
    display: block;
    width: 100%;
  }
}

.name {
  margin: 20px 0;
  font-size: 1.3rem;
  letter-spacing: 2px;;
}

.container {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 100%;
  background-color: #fafafa;
}

.textarea {
  resize: none;
  width: 80%;
  margin: 0 auto;
  display: block;
}

.images {
  width: 100%;
  display: flex;
  justify-content: space-around;
  .img {
    margin: 5%;
    width: 35%;
    img {
      display: block;
      width: 100%;
    }
  }
}

.list {
  list-style: none;
  width: 95%;
  margin: 50px auto;
}

li {
  background-color: #fff;
  border: 1px solid purple;
  padding: 5px;
  margin: 10px 0;
  border-radius: 5px;
  .text {
    margin-bottom: 10px;
    font-size: 1.1rem;
    letter-spacing: 2px;
  }
  .time {
    font-size: 0.7rem;
  }
}

.yakatabune {
  width: 70%;
  margin: 0 auto;
  img {
    display: block;
    width: 100%;
  }
}
</style>
