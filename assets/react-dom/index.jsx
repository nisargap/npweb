var Index = React.createClass({
  render: function() {

    return (
    <div>
      <h1>NpWeb</h1>
      <Menu />
      <Login />
     </div>
    );
  }
});

ReactDOM.render(
  <Index />,
  document.getElementById('index')
);