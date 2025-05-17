import { post } from "aws-amplify/api";
import "./App.css";

function App() {
  return (
    <>
      <button
        onClick={() => {
          post({
            apiName: "orderpackscalculator",
            path: "/calculate",
            options: {
              body: {
                packSizes: [23, 31, 53],
                amount: 500000,
              },
            },
          }).response.then((res) => {
            console.log(res);
          });
        }}
      >
        Calculate
      </button>
    </>
  );
}

export default App;
