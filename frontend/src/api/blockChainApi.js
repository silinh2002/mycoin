import API from './api';

const createwallet = () => {
  return API.post('/createwallet').then(res => res.data)
}

const getBalance = (params) => {
  return API.get(`/getbalance/${params}`).then(res => res.data)
}

const sendTransaction = (params) => {
  return API.post('/send', params).then(res => res.data)
}

const getHistoryAll = async () => {
  return await API.get('/histories-all').then(res => res.data)
}


// eslint-disable-next-line import/no-anonymous-default-export
export default {
  createwallet,
  getBalance,
  sendTransaction,
  getHistoryAll
}