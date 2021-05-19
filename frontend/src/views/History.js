import React, { useState, useEffect } from "react";
import Table from "react-bootstrap/Table";
import * as moment from "moment";
import httpBlockChain from "api/blockChainApi";

export default function HistoryTransaction() {
  const [transactions, setTransactions] = useState([]);
  useEffect(async () => {
    let data = await httpBlockChain.getHistoryAll();
    setTransactions(data);
  }, []);

  return (
    <div className="history-container">
      <h3>History Transaction</h3>
      <Table striped bordered hover className="table">
        <thead>
          <tr>
            <th className="cel-hash">Txn Hash</th>
            <th className="cel-age">Age</th>
            <th className="cel-from">From</th>
            <th className="cel-to">To</th>
            <th className="cel-val">Value</th>
          </tr>
        </thead>
        <tbody>
          {transactions.length > 0 &&
            transactions.map((item, i) => (
              <tr key={i}>
                <td className="cel-hash">{item.TXID}</td>
                <td className="cel-age">
                  {moment(new Date(item.Timestamp * 1000)).format(
                    "YYYY-MM-DD HH:mm"
                  )}
                </td>
                <td className="cel-from">
                  {item.From ? item.From : "system-admin"}
                </td>
                <td className="cel-to">{item.To}</td>
                <td className="cel-val">{item.Value}</td>
              </tr>
            ))}
        </tbody>
      </Table>
    </div>
  );
}
