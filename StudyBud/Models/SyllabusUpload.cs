using System.ComponentModel.DataAnnotations;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.RazorPages;

namespace StudyBud.Models;

public class SyllabusUpload : PageModel
{
    [BindProperty]
    public FileUpload FileUpload { get; set; }
}

public class FileUpload
{
    [Required]
    [Display(Name = "File")]
    public IFormFile FormFile { get; set; }
}
