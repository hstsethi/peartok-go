document.getElementById('sendForm').addEventListener('submit', function(event) {

                                          event.preventDefault();


                                          const amount = document.getElementById('amount').value;

                                          const receiver = document.getElementById('receiver').value;


                                          fetch('http://localhost:8080/send', {


                                                  method: 'POST',


                                                  headers: {


                                                                                  'Content-Type': 'application/json',


                                                                              },


                                                  body: JSON.stringify({ amount: parseFloat(amount), receiver: receiver }),


                                          })

                                          .then(response => response.json())

                                          .then(data => console.log(data))

                                          .catch((error) => {


                                                  console.error('Error:', error);


                                          });

                                      });
</script>
