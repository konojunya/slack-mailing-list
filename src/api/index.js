import axios from "axios"

class Client {

  constructor() {
    this.http = axios.create({
      baseURL: "/api"
    })
  }

  getUsers() {
    return this.http.get("/users")
  }
  postMessage(text = "", userIds = []) {
    return this.http.post("/message", {
      text,
      userIds
    })
  }

}

export default new Client