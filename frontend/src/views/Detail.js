import React, { useState, useEffect } from 'react';
import BlockInfo from './BlockInfo';
import httpBlockChain from 'api/blockChainApi'

export default function Detail() {
  const [address, setAddress] = useState();
  const [addressState, setAddressState] = useState();
  const getDetail = async () => {
    let data = await httpBlockChain.getBalance(address)
    setAddressState(data.Balance)
  }

  return (
    <div className="detail-container">
      <div className="tmp-form-group">
        <div className="tmp-form-control">
          <label className="tmp-form-control__label">My Address</label>
          <input type="text" placeholder="0xCD124" onChange={e => setAddress(e.target.value)} className="tmp-form-control__input" />
        </div>
      </div>
      <button className="detail__btn" onClick={getDetail}>Submit</button>
      <BlockInfo address={address} balance={addressState || 0}/>
    </div>
  )
}