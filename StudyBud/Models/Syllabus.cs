namespace StudyBud.Models;

public class Syllabus
{
	public string SyllabusID { get; set; }

	public string ClassTitle { get; set; }

	public string School { get; set; }

	public Professor Professor { get; set; }

	public bool TA { get; set; }

	public int CreditHours { get; set; }

	public DateTime StartDate { get; set; }

	public DateTime EndDate { get; set; }

	public string Semester { get; set; }

	public string CourseObjectives { get; set; }

	public List<Assignment> Assignments { get; set; }

	public string Objectives { get; set; }

	public List<Book> Books { get; set; }

	//grading scale?

	public string TechRequirements { get; set; }

	public string Misc { get; set; }

	public bool LateWork { get; set; }

	public byte[] Content { get; set; }

}


