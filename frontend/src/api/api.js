import axios from "axios";

const Api = axios.create({
  // baseURL: process.env.REACT_APP_BASE_URL,
  baseURL: "http://localhost:1327",
  withCredentials: false,
  headers: {
    Accept: "application/json",
    "Content-Type": "application/json",
  },
});

export default Api;
