import axios from "axios";

const apiClient = axios.create({
  headers: {
    "Content-type": "application/json",
    "Accept": "application/json"
  },
});

export default apiClient;