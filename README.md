# booking-app

 A booking application made in Go that handles the reservation of tickets for a Python Conference. The program incorporates various features such as user input validation, ticket booking, and email confirmation using concurrency.

The main() function is the entry point of the program. It begins by displaying a welcome message, informing users about the available tickets. It then enters a loop to continuously accept user input for booking tickets.The helper.ValidateUser() function is called to validate the user's input, checking if the email format is valid, the name contains only alphabetic characters, and the number of tickets is within the remaining availability. If the input is valid, the bookTickets() function is called to update the remaining ticket count, store the user's booking details in the bookings slice, and display a confirmation message. Additionally, the sendTickets() function is executed concurrently using goroutines to simulate the process of sending tickets via email.

The program continues to loop until all the tickets are booked or the user enters invalid input. If all tickets are booked, the program ends with a message indicating the unavailability of tickets.

