import axios from 'axios'

class HttpService {

  GetProfile() {
    const id = location.pathname.split("/@")[1]
    return axios.get("/api/profile", {
      params: {
        id
      }
    })
  }

  PostMessage(message) {
    const id = location.pathname.split("/@")[1]
    return axios.post(`/api/@${id}`, {
      body: message
    })
  }

  GetMessages() {
    const id = location.pathname.split("/@")[1]
    return axios.get(`/api/@${id}`)
  }

}

export default new HttpService();