import React, { useEffect } from "react";
import reactLogo from "./assets/react.svg";
import viteLogo from "/vite.svg";
import "./App.css";
import Bookshelf from "./components/Bookshelf";

const App: React.FC = () => {
  useEffect(() => {
    const checkHealth = async (retries = 5) => {
      for (let i = 0; i < retries; i++) {
        try {
          // FIXME:URLのハードコードを避ける
          const response = await fetch("http://localhost:8080/health");
          if (response.ok) {
            const data = await response.json();
            console.log("Health check: ", data);
            return;
          } else {
            console.error("Health check failed: ", response.status);
          }
        } catch (error) {
          console.error("Error during health check: ", error);
        }
        console.log(`Retrying health check... (${i + 1}/${retries})`);
        await new Promise((resolve) => setTimeout(resolve, 5000)); // 5秒待機
      }
    };

    checkHealth();
  }, []);

  return (
    <>
      <div className="flex justify-center">
        <a href="https://vite.dev" target="_blank">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
        <a href="https://react.dev" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>

      <a
        href="https://books.google.co.jp/books"
        target="_blank"
        rel="noopener noreferrer"
        className="text-2xl"
      >
        Googleブックを開く
      </a>

      <Bookshelf />
    </>
  );
};

export default App;
