import React, { useState, useEffect } from 'react';
import httpBlockChain from 'api/blockChainApi'


export default function SendTransaction() {
  const [sender, setSender] = useState()
  const [recipient, setRecipient] = useState()
  const [amount, setAmount] = useState(0)

  const onSendTransaction = async () => {
    let data = await httpBlockChain.sendTransaction({ from: sender, to: recipient, amount: amount })
    alert(data.Result)
  }

  return (
    <div className="transaction-container">
      <div className="transaction">
        <div className="transaction__header"><span>Send Transaction</span></div>
        <div className="transaction__body">
          <div className="transaction__form">
            <div className="transaction__row tmp-form-control">
              <label className="tmp-form-control__label">From</label>
              <input onChange={e => setSender(e.target.value)} type="text" className="tmp-form-control__input" />
            </div>
            <div className="transaction__row group-double">
              <div className="tmp-form-control group-double__2">
                <label className="tmp-form-control__label">Amount</label>
                <input value={amount} onChange={e => setAmount(+e.target.value)} type="number" placeholder="" className="tmp-form-control__input" />
              </div>
            </div>
            <div className="transaction__row tmp-form-control">
              <label className="tmp-form-control__label">To</label>
              <input onChange={e => setRecipient(e.target.value)} type="text" className="tmp-form-control__input" />
            </div>
          </div>
        </div>
        <div className="transaction__footer">
          <button onClick={onSendTransaction} className="btn btn-disabled">Send Transaction</button>
        </div>
      </div>
    </div>
  )
}