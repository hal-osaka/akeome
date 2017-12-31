<template>
  <section>

    <div class="yakatabune">
      <img src="https://s3-us-west-2.amazonaws.com/akeome/yakatabune.png" alt="yakatabune">
    </div>

    <h1>お正月メッセージを送り合おう！</h1>

    <div class="container">
      <ol>
        <li>メッセージをしたい人のTwitterIDを入力してください！</li>
        <li>メッセージを送信してください！</li>
        <li>いろんな人のメッセージを覗き見してみましょう！</li>
      </ol>

      <div class="input-div">
        <b-field horizontal>
          <b-input placeholder="TwitterID" expanded v-model="twitterid"></b-input>
        </b-field>
        <b-field horizontal>
          <p class="control">
            <button class="button is-primary" @click="search">探す</button>
          </p>
        </b-field>
      </div>

    </div>

    <div class="images">
      <div class="img">
        <img src="https://s3-us-west-2.amazonaws.com/akeome/hime.png" alt="hime">
      </div>
      <div class="img">
        <img src="https://s3-us-west-2.amazonaws.com/akeome/tono.png" alt="tono">
      </div>
    </div>

  </section>
</template>

<script>
export default {
  name: "howto",
  data() {
    return {
      twitterid: ""
    }
  },
  methods: {
    search() {
      location.pathname = `/@${this.twitterid}`
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

.yakatabune {
  width: 70%;
  margin: 0 auto;
  img {
    display: block;
    width: 100%;
  }
}

.container {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 100%;
  background-color: #fafafa;
}

ol {
  font-size: 0.9rem;
}

li {
  margin: 20px 30px;
}

.images {
  width: 100%;
  display: flex;
  justify-content: space-around;
  margin: 30px 0;
  padding: 0 30px;
  .img {
    margin: 5%;
    width: 35%;
    img {
      display: block;
      width: 100%;
    }
  }
}
</style>
