import axios from "axios"

class Client {

  constructor() {
    this.http = axios.create({
      baseURL: "/api"
    })
  }

  getUsers(nextCursor) {
    return this.http.get(`/users?next_cursor=${nextCursor}`)
  }
  
  getChannels(nextCursor) {
    return this.http.get(`/channels?next_cursor=${nextCursor}`)
  }

  postMessage(text = "", userIds = []) {
    return this.http.post("/message", {
      text,
      ids: userIds
    })
  }

}

export default new Client