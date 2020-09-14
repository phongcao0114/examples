1. Run book-manager-service
2. Run gateway
3. Access: localhost:9876
4. Define query to get list of books:

	query {
	  books{
	    id,
	    name,
	    author,
	  }
	}
