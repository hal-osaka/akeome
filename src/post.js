import Vue from 'vue'
import Buefy from 'buefy'
import 'buefy/lib/buefy.css'

import Post from './components/post/index.vue'

Vue.use(Buefy)
Vue.component("post", Post)

new Vue({
  el: "#app"
})