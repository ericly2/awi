import { useCallback, useEffect, useState } from "react";

const App = () => {
  const [book, setBook] = useState(null);

  const getBookAPI = useCallback(async () => {
    let response = await fetch('/book')
    response = await response.json();
    setBook(response);
  }, [])

  useEffect(() => {
    getBookAPI()
  }, [getBookAPI])

  return (
    <div>
      <h1>This is a test</h1>
      {book != null ? 
        <div>
          <div>{book.title}</div>
          <div>{book.author}</div>
          <div>{book.pages}</div>
        </div>
        : null}
    </div>
  );
};

export default App;
