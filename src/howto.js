import Vue from 'vue'
import Buefy from 'buefy'
import 'buefy/lib/buefy.css'

import HowTo from './components/index.vue'

Vue.use(Buefy)
Vue.component("howto", HowTo)

new Vue({
  el: "#app"
})