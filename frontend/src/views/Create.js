import React, { useState, useEffect } from 'react';
import 'styles/create.scss';
import icLogin from 'assets/login.svg';
import httpBlockChain from 'api/blockChainApi'

export default function Register() {
  const [address, setAddress] = useState()
  const register = async () => {
    // dispatch(registerWallet())
    let data = await httpBlockChain.createwallet()
    setAddress(data.Address)
  }
  // let address = useSelector(state => state.blockchain.wallet)
  return (
    <div className="login">
      <div className="login__wrapper">
        <div className="login__icon"><img src={icLogin} alt="" /></div>
        <div className="login__title">Your Mnemonic Phrase</div>
        <div className="login__footer">
          <button onClick={register} className="login__btn">Create wallet</button>
        </div>
        { address && <div className="login__wallet">
          <span>My address create</span>
          <p>{address}</p>
        </div>}
      </div>
    </div>
  )
}