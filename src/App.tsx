import { post } from "aws-amplify/api";
import "./App.css";
import { useState } from "react";

function App() {
  const [packSizes, setPackSizes] = useState<string>("");
  const [amount, setAmount] = useState<string>("");
  const [orderPacks, setOrderPacks] = useState<{ [key: number]: number }>({});

  return (
    <>
      <button
        onClick={async () => {
          const { body } = await post({
            apiName: "orderpackscalculator",
            path: "/calculate",
            options: {
              body: {
                packSizes: packSizes.split(",").map(Number),
                amount: Number(amount),
              },
            },
          }).response;

          const calculateResponse = (await body.json()) as {
            orderPacks: { [key: number]: number };
          };

          setOrderPacks(calculateResponse.orderPacks);
        }}
      >
        Calculate
      </button>

      <div>
        <label htmlFor="packSizes">Pack Sizes (comma separated)</label>
        <input
          id="packSizes"
          type="text"
          value={packSizes}
          onChange={(e) => setPackSizes(e.target.value)}
        />
      </div>

      <div>
        <label htmlFor="amount">Amount (integer)</label>
        <input
          id="amount"
          type="text"
          value={amount}
          onChange={(e) => setAmount(e.target.value)}
        />
      </div>

      <div>
        {Object.entries(orderPacks).map(([packSize, quantity]) => (
          <div key={packSize}>
            {packSize} x {quantity}
          </div>
        ))}
      </div>
    </>
  );
}

export default App;
