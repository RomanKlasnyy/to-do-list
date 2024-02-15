import org.hibernate.Session;
import org.hibernate.SessionFactory;
import org.hibernate.cfg.Configuration;
import javax.persistence.*;
import java.util.Date;
import java.util.List;
import java.util.Scanner;

@Entity
@Table(name = "task")
class ToDoList {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private int id;

    @Column(name = "task")
    private String task;

    @Column(name = "deadline")
    private Date deadline;

    public ToDoList() {}

    public ToDoList(String task, Date deadline) {
        this.task = task;
        this.deadline = deadline;
    }

    public String getTask() {
        return task;
    }

    public Date getDeadline() {
        return deadline;
    }
}

public class Main {

    public static void main(String[] args) {
        SessionFactory factory = new Configuration()
                .configure("hibernate.cfg.xml")
                .addAnnotatedClass(ToDoList.class)
                .buildSessionFactory();

        Scanner scanner = new Scanner(System.in);
        while (true) {
            try (Session session = factory.getCurrentSession()) {
                session.beginTransaction();
                System.out.println("1) Today's tasks\n" +
                        "2) Week's tasks\n" +
                        "3) All tasks\n" +
                        "4) Missed tasks\n" +
                        "5) Add task\n" +
                        "6) Delete task\n" +
                        "0) Exit");
                String choice = scanner.nextLine();

                switch (choice) {
                    case "0":
                        System.exit(0);
                        break;
                    case "1":
                        // Today's tasks
                        break;
                    case "2":
                        // Week's tasks
                        break;
                    case "3":
                        // All tasks
                        break;
                    case "4":
                        // Missed tasks
                        break;
                    case "5":
                        // Add task
                        break;
                    case "6":
                        // Delete task
                        break;
                    default:
                        System.out.println("Invalid choice. Please try again.");
                }
                session.getTransaction().commit();
            } catch (Exception e) {
                System.out.println("Error");
            }
        }
    }
}
