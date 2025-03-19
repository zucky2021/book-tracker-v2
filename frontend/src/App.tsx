import React from "react";
import reactLogo from "./assets/react.svg";
import viteLogo from "/vite.svg";
import "./App.css";
import Bookshelf from "./components/Bookshelf";

const App: React.FC = () => {
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
        href="https://books.google.co.jp/books?uid=100173087971504642758&hl=ja"
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
