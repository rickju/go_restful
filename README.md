Requirements
--------------------

You will be implementing a simple API. You may use any language or technology
you prefer. It's recommended that you set up a working development environment
ahead of time.

You will have internet access and are allowed to reference documentation or
search for help on Google, Stack Overflow, or anywhere else.  Email
evan@vendhq.com with any questions you have.

You should have been given access to a clean github repository in advance of
the test. Please commit your code (at least) as you finish each stage.

Please take a  few minutes to read the test and ask questions. Once you start coding you have 2.5 hours.
Don’t worry if you don’t complete the entire test.
We’re more interested in how you work and what you think about than whether you complete the entire test. 

Stage 1 GET /products
------------------
Create an API endpoint that serves a list of products that can be sold. 

A product consists of:
  A unique ID
  A name
  A price

Products may be stored in memory. An external database may be used, but is not required.

  ID Name Price
  1 Chrome Toaster $100
  2 Copper Kettle $49.99
  3 Mixing Bowl $20


Stage 2 POST /products
------------------
  Create an API endpoint that allows creating a new product. 
  The response should contain the product details in the same format as Stage 1.
  The new product must be persisted so that the endpoint from stage 1 includes the new product in its responses.

Stage 3 POST /sales
------------------
  Create an API endpoint that allows sales to be made. It is not necessary to persist sales.

  A sale request consists of:
    An array of line items.
    Each line item includes a product ID and a quantity.

  A sale response consists of:
    Everything in the sale request
    A total for each line item.
    A total cost of sale.

Stage 4
------------------

  Modify the sales API to allow discounts on the overall sale. 
  The discount on the sale is a flat dollar amount (i.e. $10 off the total as opposed to 10% off the total).

  The sale request is modified to add a new discount field, representing the total discount across the entire sale.

  For tax reasons, the discount must be spread proportionally across the line items in the sale.
  Each line item in the response should also include a discount field containing the proportion of the discount for that line item.


