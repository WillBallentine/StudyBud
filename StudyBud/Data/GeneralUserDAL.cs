using System;
using Microsoft.EntityFrameworkCore;
using StudyBud.Data.Interfaces;
using StudyBud.Models;


namespace StudyBud.Data
{
	public class GeneralUserDAL : IGeneralUserDAL
	{
		ApplicationDbContext _dbContext;

		public GeneralUserDAL(ApplicationDbContext dbContext)
		{
			_dbContext = dbContext;
		}

	}
}

