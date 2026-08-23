namespace FlatData;

public enum TestStatus : int
{
	None = 0,
	Ready = 1
}
public struct TestSchema : IFlatbufferObject
{
	public int Id { get; }
	public int ItemsLength { get; }
	public float WeightLength { get; }

	public int Items(int index)
	{
		return 0;
	}

	public float Weight(int index)
	{
		return 0;
	}

	public static void AddId(FlatBufferBuilder builder, int id)
	{
	}

	public static void StartItemsVector(FlatBufferBuilder builder, int count)
	{
	}

	public static void CreateItemsVector(FlatBufferBuilder builder, int[] values)
	{
	}

	public static void StartWeightVector(FlatBufferBuilder builder, int count)
	{
	}

	public static void CreateWeightVector(FlatBufferBuilder builder, float[] values)
	{
	}

	public static void FinishTestSchemaBuffer(FlatBufferBuilder builder, int offset)
	{
	}
}
